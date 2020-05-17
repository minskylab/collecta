package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/facebookincubator/ent/entc"
	"github.com/facebookincubator/ent/entc/gen"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/minskylab/collecta/errors"
)

func generateGQLGenFromEntSchema(entSchemaPath string, gqlgenAPIPath string) {
	graph, err := entc.LoadGraph(entSchemaPath, &gen.Config{})
	if err != nil {
		panic(errors.Wrap(err, "error at load graph"))
	}

	for _, sch := range graph.Schemas {
		depsBuffer := bytes.NewBufferString("")
		schBuffer := bytes.NewBufferString("")
		if _, err = fmt.Fprintf(schBuffer,"type %s {\n", sch.Name); err != nil {
			panic(errors.Wrap(err, "error at write buffer"))
		}

		for _, f := range sch.Fields {
			if f.Info.Type == field.TypeEnum {
				// to enum buffer
				// enumName := strings.TrimPrefix(sch.Name, strings.ToLower(sch.Name) + ".")

				enumName := strings.ToUpper(string(f.Name[0])) + f.Name[1:]
				enumName = sch.Name + enumName

				if _, err = fmt.Fprintf(depsBuffer,"enum %s {\n", enumName); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}

				for _, val := range f.Enums {
					if _, err = fmt.Fprintf(depsBuffer,"    %s\n", val); err != nil {
						panic(errors.Wrap(err, "error at write buffer"))
					}
				}

				if _, err = fmt.Fprint(depsBuffer,"}"); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}

				// to schema buffer
				if _, err = fmt.Fprintf(schBuffer,"    %s: ", f.Name); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}

				if _, err = fmt.Fprintf(schBuffer,"%s", enumName); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}

				if !f.Optional {
					if _, err = fmt.Fprint(schBuffer,"!"); err != nil {
						panic(errors.Wrap(err, "error at write buffer"))
					}
				}

			} else {
				// to schema buffer
				if _, err = fmt.Fprintf(schBuffer,"    %s: ", f.Name); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}
				gqlType, err := gqlTypeFromTypeInfo(f.Info)
				if err != nil {
					panic(errors.Wrap(err, "error at try to parse the gql type"))
				}
				if _, err = fmt.Fprintf(schBuffer,"%s", gqlType); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}

				if !f.Optional {
					if _, err = fmt.Fprint(schBuffer,"!"); err != nil {
						panic(errors.Wrap(err, "error at write buffer"))
					}
				}
			}

			if _, err = fmt.Fprint(schBuffer,"\n"); err != nil {
				panic(errors.Wrap(err, "error at write buffer"))
			}

		}


		for _, e := range sch.Edges {
			if _, err = fmt.Fprintf(schBuffer,"    %s: ", e.Name); err != nil {
				panic(errors.Wrap(err, "error at write buffer"))
			}

			if !e.Unique {
				if _, err = fmt.Fprint(schBuffer,"["); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}
			}

			if _, err = fmt.Fprint(schBuffer,e.Type); err != nil {
				panic(errors.Wrap(err, "error at write buffer"))
			}

			if !e.Required {
				if _, err = fmt.Fprint(schBuffer,"!"); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}
			}

			if !e.Unique {
				if _, err = fmt.Fprint(schBuffer,"]!"); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}
			}

			if _, err = fmt.Fprint(schBuffer, "\n"); err != nil {
				panic(errors.Wrap(err, "error at write buffer"))
			}
		}

		if _, err = fmt.Fprint(schBuffer,"}"); err != nil {
			panic(errors.Wrap(err, "error at write buffer"))
		}

		filename := strings.ToLower(sch.Name) + ".graphqls"
		schemaFilepath := path.Join(gqlgenAPIPath, "graph", "schema")

		_ = os.MkdirAll(schemaFilepath, os.ModePerm)

		finalSchemaFilename := path.Join(schemaFilepath, filename)

		data := []byte(schBuffer.String())
		if depsBuffer.String() != "" {
			data = []byte(depsBuffer.String() + "\n\n" + string(data))
		}

		if err := ioutil.WriteFile(finalSchemaFilename, data, 0644); err != nil {
			panic(errors.Wrap(err, "error at try to write custom graphQL schema"))
		}

	}

	gqlConfig := path.Join(gqlgenAPIPath, "gqlgen.yml")
	data, err := ioutil.ReadFile(gqlConfig)
	if err != nil {
		panic(errors.Wrap(err, "error at try to read gql config file"))
	}

	indent := 1

	yamlConfigBuffer := bytes.NewBufferString("")
	if _, err = fmt.Fprint(yamlConfigBuffer,"models:\n"); err != nil {
		panic(errors.Wrap(err, "error at write buffer"))
	}

	id := strings.Repeat("  ", indent) + "ID:\n" +
		strings.Repeat("  ", indent+1) + "model:\n" +
		strings.Repeat("  ", indent+2) + "- github.com/google/uuid.UUID\n"

	if _, err = fmt.Fprint(yamlConfigBuffer, id); err != nil {
		panic(errors.Wrap(err, "error at write buffer"))
	}

	for _, sch := range graph.Schemas {
		for _, f := range sch.Fields { // Validate if it has a enum field
			if f.Info.Type == field.TypeEnum {
				// to enum buffer
				// enumName := strings.TrimPrefix(sch.Name, strings.ToLower(sch.Name) + ".")

				rawEnumName := strings.ToUpper(string(f.Name[0])) + f.Name[1:]
				enumName := sch.Name + rawEnumName

				if _, err = fmt.Fprintf(yamlConfigBuffer,"%s%s:\n", strings.Repeat("  ", indent), enumName); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}

				indent += 1

				if _, err = fmt.Fprintf(yamlConfigBuffer,"%s%s:\n", strings.Repeat("  ", indent), "model"); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}

				indent += 1

				rootPackagePath := "github.com/minskylab/collecta/ent/" + strings.ToLower(sch.Name) + "." + rawEnumName
				if _, err = fmt.Fprintf(yamlConfigBuffer,"%s%s\n", strings.Repeat("  ", indent), "- "+rootPackagePath); err != nil {
					panic(errors.Wrap(err, "error at write buffer"))
				}

				indent -= 2
			}
		}

		if _, err = fmt.Fprintf(yamlConfigBuffer,"%s%s:\n", strings.Repeat("  ", indent), sch.Name); err != nil {
			panic(errors.Wrap(err, "error at write buffer"))
		}

		indent += 1

		if _, err = fmt.Fprintf(yamlConfigBuffer,"%s%s:\n", strings.Repeat("  ", indent), "model"); err != nil {
			panic(errors.Wrap(err, "error at write buffer"))
		}

		indent += 1

		rootPackagePath := "github.com/minskylab/collecta/ent." + sch.Name

		if _, err = fmt.Fprintf(yamlConfigBuffer,"%s%s\n", strings.Repeat("  ", indent), "- "+rootPackagePath); err != nil {
			panic(errors.Wrap(err, "error at write buffer"))
		}

		indent -= 1

		if _, err = fmt.Fprintf(yamlConfigBuffer,"%s%s:\n", strings.Repeat("  ", indent), "fields"); err != nil {
			panic(errors.Wrap(err, "error at write buffer"))
		}

		indent += 1

		for _, e := range sch.Edges {
			if _, err = fmt.Fprintf(yamlConfigBuffer,"%s%s:\n", strings.Repeat("  ", indent), e.Name); err != nil {
				panic(errors.Wrap(err, "error at write buffer"))
			}

			indent += 1

			if _, err = fmt.Fprintf(yamlConfigBuffer,"%sresolver: true\n", strings.Repeat("  ", indent)); err != nil {
				panic(errors.Wrap(err, "error at write buffer"))
			}

			indent -= 1
		}

		// if _, err = fmt.Fprint(yamlConfigBuffer,"\n"); err != nil {
		// 	panic(errors.Wrap(err, "error at write buffer"))
		// }
		indent -= 2
	}

	lines := strings.Split(string(data), "\n")
	modelStartCutPoint := -1
	modelEndCutPoint := -1
	for i, l := range lines {
		if strings.Contains(l, "models:") {
			modelStartCutPoint = i
		}
		if modelStartCutPoint != -1 && strings.TrimPrefix(l," ") == l {
			modelEndCutPoint = i - 1
		}
	}

	lines = append(lines[:modelStartCutPoint], lines[modelEndCutPoint:]...)
	lines[modelStartCutPoint] = yamlConfigBuffer.String()

	insertedConfigYaml := strings.Join(lines, "\n")

	finalGqlGenConfigPath := path.Join(gqlgenAPIPath, "gqlgen.yml")
	if err := ioutil.WriteFile(finalGqlGenConfigPath, []byte(insertedConfigYaml), 0644); err != nil {
		panic(errors.Wrap(err, "error at try to write new gql config file"))
	}
}
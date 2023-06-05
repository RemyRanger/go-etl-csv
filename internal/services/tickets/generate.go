package tickets

//go:generate oapi-codegen --old-config-style -o handler/api_types.gen.go -package=handler -include-tags=Tickets -generate types ../../../doc/dist/api.yaml
//go:generate oapi-codegen --old-config-style -o handler/api.gen.go -package=handler -include-tags=Tickets -generate chi-server ../../../doc/dist/api.yaml

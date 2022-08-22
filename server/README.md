TODO:

 - [ ] adicionar serviço como serveless functions? separar API em diferentes funções
 - [ ] adicionar pprof ou prometheus
 - [ ] .gitignore config.yaml
 - [ ] adicionar mais coisas no logger?
 - [ ] melhorar respostas da API com echo.NewHTTPError(http.StatusBadRequest, "mensagem")
 - [ ] arrumar testes de inventory (criar o mock de service separado e importar?)
 - [ ] melhorar os testes
 - [ ] echo.Bind despadronizado, alguns InternalServer e outro BadRequest (padronizar)
 - [ ] middleware de binding?
 - [ ] validation (email, password, required, etc)
 - [ ] criar um error no echo para centralizar o log
 - [ ] deletar repo.go?
 - [ ] utilizar echo.ErrUnauthorized .ErrBadRequest etc?
 - [ ] adicionar salt no password
 - [ ] middleware de binding
 - [ ] adicionar preço por unidade (price unit_price quantity_price)
 - [ ] adicionar campo para imagens nas entidades? (product), (inventory)
 - [ ] adicionar logger
 - [ ] search by name
 - [ ] search customers -> first name para todas colunas
 - [ ] padronizar as repostas de erro ao usuário da api
 - [ ] adicionar contratos/inventorios na entidade user?

reference:
 - [ddd-go-template by @VinGracia](https://github.com/VinGarcia/ddd-go-template/blob/master/v2-domain-adapters-and-helpers)
 - [ddd-hexagonal-onion-clean](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/)
 - [applying-the-clean-architecture-to-go-applications](https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)
 - [Elton Minetto](https://eltonminetto.dev/post/2020-06-29-clean-architecture-2anos-depois/)


 {base-path}/{area}/{version}/entity1/{entity1}/{entity2}
 base-path = {dns-name}/{microservice-name}
 area = api or management to indicate general area (just use api)

TODO:

 - [ ] adicionar logger
 - [ ] search by name
 - [ ] separar entidade users -> produtores, compradores
 - [ ] search customers -> first name para todas colunas
 - [ ] unique constraints no DB, validação
 - [ ] padronizar as repostas de erro ao usuário da api

reference:
 - [ddd-go-template by @VinGracia](https://github.com/VinGarcia/ddd-go-template/blob/master/v2-domain-adapters-and-helpers)
 - [ddd-hexagonal-onion-clean](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/)
 - [applying-the-clean-architecture-to-go-applications](https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)
 - [Elton Minetto](https://eltonminetto.dev/post/2020-06-29-clean-architecture-2anos-depois/)


 {base-path}/{area}/{version}/entity1/{entity1}/{entity2}
 base-path = {dns-name}/{microservice-name}
 area = api or management to indicate general area (just use api)

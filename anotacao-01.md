## Api Products

Basicamente um usuario pode criar uma conta e cadastrar seus produtos e categorias para seus produtos.
Também teremos uma timeline listando todos os produtos e seus usuários com filtro por: "Nome","Categoria", "Preço".
Ao clicar no produto dentro da timeline sera listado todos os detalhes do produto.

Para poder comprar o produto o usuario deve ser cadastrar na plataforma, e adicionar o carrinho de compras

### User

- [x] Um usuário pode se registrar na api
- [x] Um usuário pode logar-se na API através de jwt token

### Categoria dos produtos do usuario

- [x] O usuário pode cadastrar categorias para seus produtos
- [] O usuário pode deletar categorias
- [] O usuário pode atualizar o nome da categoria
- [] O usuário pode listar todas as suas categorias
- [] O usuário pode listar todas as suas categorias juntamente com os produtos de cada cetegoria
- [] O usuário pode pesquisar categoria pelo nome
- [] Para cada categoria selecionada deve-se trazer os produtos referentes a ela
- [] Ao deletar uma categoria deve-se deletar os produtos referente a ela
  juntamente com todas suas imagens

#### Produtos do usuário

- [] O usuário pode cadastrar um produto referenciando-o a uma categoria
- [] O usuario pode adicionar imagens a esse produto
- [] O usuário pode pode deletar seus produtos
- [] O usuário pode atualizar informações do seu produto
- [] O usuário pode alterar as imagens do produto
- [] Para as imagens do produto a api deve renderizar o tamanho das imagens caso ele seja maior que 1024px para 800px
- [] O usuário pode listar seus produtos
- [] O usuário pode pesquisar pode nome, descrição e preço seus produtos

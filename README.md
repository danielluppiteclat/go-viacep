# go-viacep
Consulta de CEPs usando a API ViaCEP em Go

A aplicação cria um struct chamado Endereco, que define os campos que o ViaCEP retorna e extrai o parâmetro CEP da query usada para teste.

Se o CEP não for informado, na query de teste, é retornado um erro para o usuário.

Quando o CEP é informado, a aplicação monta a URL do ViaCEP e faz a requisição, lê a resposta e a converte de JSON para a struct Endereco, retornando a resposta, definida como JSON.

A aplicação converte a resposta do viaCep para a estrutura Endereo para restringir o retorno, porque o ViaCEP retorna várias informações adicionais, por exemplo IBGE, GIA, DDD, SIAFI, etc., que não são relevantes para o objetivo da aplicação.

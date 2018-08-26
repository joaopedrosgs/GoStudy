# Pagina inicial com agenda de provas
Este programa faz uso da agenda do Google para obter a data das provas, a ideia é que seja iniciado com o sistema e utilizado como pagina inicial do navegador
![Alt text](print.png?raw=true "Print")

# Como utilizar:

1. Gere um token de acesso do Google (https://console.developers.google.com/apis/dashboard) e coloque-o na pasta em um arquivo chamado token.json da seguinte forma
 
 `{"access_token":"token","expiry":"data"}` 

2. Crie um arquivo de sites chamado sites.json no seguinte formato:
  
  `{`

       "sites": [

           "site1",

           "site2"

       `]`
   `}`



   Seja feliz (ou triste, já que tem muita prova)

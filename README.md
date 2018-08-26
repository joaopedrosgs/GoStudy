# GoStudy
Este programa faz uso da agenda do Google para obter a data das provas, a ideia é que seja iniciado com o sistema e utilizado como pagina inicial do navegador
![Alt text](print.png?raw=true "Print")

# Como utilizar:

1. Gere um token de acesso do Google (https://console.developers.google.com/apis/dashboard) e coloque-o na pasta "$HOME/.config/goStudy" em um arquivo chamado token.json da seguinte forma
 
 `{"access_token":"token","expiry":"data"}` 

2. crie o websites.json da pasta "$HOME/.config/goStudy" para adicionar seus sites

   Seja feliz (ou triste, já que tem muita prova)

# Exemplo de websites.json
```json
[
    {
        "Icon": "https://s.glbimg.com/en/ho/static/globo_com_2016/img/home_200x200.png",
        "Name": "Globo",
        "Url": "https://www.globo.com",
        "Description": "Globo.com"
    },
    {
        "Icon": "http://icons.iconarchive.com/icons/marcus-roberto/google-play/256/Google-Drive-icon.png",
        "Name": "GDrive PET",
        "Url": "https://drive.google.com/drive/folders/0B8Gfc8vlUUjsN0xEV1otaFpvZXM",
        "Description": "Drive do PET"
    },
    {
        "Icon": "https://hp.imguol.com.br/c/home/interacao/facebook/logo-uol.png",
        "Name": "UOL",
        "Url": "https://www.uol.com.br",
        "Description": "Uol.com.br"
    }
]
```

# TODO
-   Adicionar opção de remover e adicionar sites de forma facil
-   Criar uma forma de configuração viavel
-   Tabs realmente listando trabalhos e aulas normais

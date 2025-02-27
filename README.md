## 🌟 Visão Geral

O **CryptoPrice API** é uma solução prática para quem precisa monitorar os preços de criptomoedas em tempo real.  
  
✅ **Busca dados em tempo real:** Consome a API da CoinGecko para obter preços atualizados de criptomoedas como Bitcoin, Ethereum e Dogecoin.  
✅ **Salva os dados em JSON:** Facilita análises futuras ou integrações com outras aplicações.  
✅ **Exposição via API REST:** Permite que outros sistemas consumam os dados de forma simples e eficiente.

---

## 🚀 2️⃣ Visão Geral do Sistema

O projeto é construído sobre três pilares fundamentais:

### 📂 1. Coleta de Dados
- **Requisições HTTP:** Utiliza o pacote `net/http` para consumir a API da CoinGecko.
- **Processamento JSON:** Decodifica os dados recebidos e os transforma em informações úteis.

### 🛠 2. Armazenamento Local
- **Salva em JSON:** Grava os dados obtidos em um arquivo `cryptos.json` para consulta posterior.

### 🌐 3. Exposição via API
- **Endpoint REST:** Cria um servidor HTTP que disponibiliza os dados coletados para outros sistemas via endpoint `/cryptos`.

---

## 🎯 3️⃣ Funcionalidades do Sistema

### 🔍 Coleta e Exibição de Dados
- Busca os preços de criptomoedas em tempo real.
- Exibe os preços no terminal com formatação amigável.

### 💾 Armazenamento Persistente
- Salva os dados em um arquivo JSON com formatação legível.

### 🌐 API REST
- Exposição dos dados através de um endpoint que retorna o JSON com informações atualizadas.

---

## 🚀 4️⃣ Tecnologias Utilizadas

<div align="center">
  <h3>Tecnologias</h3>
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="50" alt="Golang" /> &nbsp;
  <img src="https://img.shields.io/badge/HTTP-net%2Fhttp-007ACC?style=for-the-badge&logo=http&logoColor=white" height="50" alt="net/http" /> &nbsp;
  <img src="https://img.shields.io/badge/JSON-encoding%2Fjson-4B8BBE?style=for-the-badge&logo=json&logoColor=white" height="50" alt="encoding/json" />
</div>

---

## 🖥️ Getting Started

### 🔧 Pré-requisitos
Antes de começar, certifique-se de ter instalado:
- **Go (1.16+):** [Download Go](https://go.dev/dl/)

### 📥 Instalação
1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/CryptoPriceAPI.git
   cd CryptoPriceAPI

2 . **Execute o projeto:**
   ```bash
    go run main.go
   ```

3. **Teste a API:**
   
Acesse http://localhost:8080/cryptos via navegador ou ferramenta de API (Postman, cURL).

---

## 📜 **Licença**

🔓 Este projeto está licenciado sob a **[MIT License](LICENSE)**, permitindo o uso, modificação e distribuição sob os termos da licença MIT.  
 
---

## Contato

Caso tenha dúvidas, sugestões ou queira contribuir, entre em contato:

<div align="center">
  <a href="https://www.linkedin.com/in/marcoauréliomacedoprado" target="_blank">
    <img src="https://img.shields.io/static/v1?message=LinkedIn&logo=linkedin&label=&color=0077B5&logoColor=white&labelColor=&style=plastic" height="36" alt="linkedin logo" />
  </a>
  <a href="https://www.instagram.com/prado.marco1/" target="_blank">
    <img src="https://img.shields.io/static/v1?message=Instagram&logo=instagram&label=&color=E4405F&logoColor=white&labelColor=&style=plastic" height="36" alt="instagram logo" />
  </a>
  <a href="https://wa.me/5567996893356" target="_blank">
  <img src="https://img.shields.io/static/v1?message=Whatsapp&logo=whatsapp&label=&color=25D366&logoColor=white&labelColor=&style=plastic" height="36" alt="whatsapp logo" />
</a>
  <a href="https://discord.com/users/yourdiscordid" target="_blank">
    <img src="https://img.shields.io/static/v1?message=Discord&logo=discord&label=&color=7289DA&logoColor=white&labelColor=&style=plastic" height="36" alt="discord logo" />
  </a>
</div>

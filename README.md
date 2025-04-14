## wallet-app

> Need visit links for install

```bash
# postgresql
https://www.postgresql.org/download/

# golang
https://go.dev/doc/install

# docker
https://docs.docker.com/desktop/

```

---

> Copy this repository

```bash
git clone https://github.com/YuranIgnatenko/wallet-app
```

---

> Settings application

```bash
# Create new file `config.env`
mkdir config.env

# Write to file `config.env`
DATABASE_URL=postgres://user:password@address:post/wallets?sslmode=disable
```

---

> Launch app (start server)

```bash
# move root directory application
cd wallet-app

# compile project
go build cmd/main.go

# launch application
./main
```

> Testing application

```bash
./test.bat
# visit link
http://127.0.0.1:5500/coverage.html
```

![demo](/internal/static/cover.png)

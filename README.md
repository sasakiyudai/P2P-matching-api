**使い方**
```
$ git clone https://github.com/sasakiyudai/P2P-matching-api
$ make
root@c1b0383693ae:/app# go run main.go
```

**使用技術**
データベース：MySQL
認証：JWT
APIサーバー：Gin

**API**
*- GET localhost:8080/api/user*
ユーザー一覧を返す

*- GET localhost:8080/api/product*
商品一覧を返す

*- POST localhost:8080/api/auth/product/new*
新しい商品を作成
```
Header
|Authorization | Bearer [token]|
Body
{
    "name":"apple",
    "comment":"美味しい",
    "price":100
}
```

*- POST localhost:8080/api/auth/product/buy/[product_id]*
購入処理
```
Header
|Authorization | Bearer [token]|
```

*- POST localhost:8080/signup*
ユーザー登録
```
Body
{
    "name":"syudai",
    "email":"yudai14142@gmail.com",
    "password":"password"
}
```

*- POST localhost:8080/login*
ログイン　アクセストークンを得る
```
Body
{
    "name":"syudai",
    "password":"password"
}
```

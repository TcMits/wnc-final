# wnc-final

## Database
-   Postgres.
-   Triển khai trong docker compose
-   Tạo dữ liệu tự động khi chạy lệnh ```docker compose -f production.yml up```
## Quick start
Lưu ý: Để sử dụng được chức năng nhận OTP qua email, vui lòng điền email của bạn vào file config.yml (trong folder config) ở mục mail  và bạn có thể dùng email này để đăng nhập (với password mặc định là: 123456789), nếu không điền email bạn vẫn có thể dùng tài khoản khách hàng (tham khảo bên dưới), với tài khoản này sẽ có demo 8 giao dịch chuyển và nhận tiền.

Data:
-   Customer account:
    - Username: iamcustomer
    - Password: 123456789
    - Account number: 22223333444455556
    - Account number của khách hàng khác cùng ngân hàng (dùng để test): 11112222333344445
    - Account number của khách hàng khác khác ngân hàng (dùng để test): 33334444555566667
-   Employee account:
    - Username: iamemployee
    - Password: 123456789
-   Admin account:
    - Username: iamadmin
    - Password: 123456789

Run app:
```sh
# Run backend app
docker compose -f production.yml up
```
- Đợi đến khi terminal hiện ra thông báo sau là đã chạy backend thành công:
```sh
golang_sacombank       | Iris Version: 12.2.0-beta7
golang_sacombank       | 
golang_sacombank       | Now listening on: http://localhost:8080
golang_sacombank       | Application started. Press CTRL+C to shut down.
golang_tpbank          | Iris Version: 12.2.0-beta7
golang_tpbank          | 
golang_tpbank          | Now listening on: http://localhost:8081
golang_tpbank          | Application started. Press CTRL+C to shut down.
```

Generate docs customeSacombankr app files:
```sh
swag init --exclude ./internal/controller/http/v1/services/employee,./internal/controller/http/v1/services/admin,./internal/controller/http/v1/services/partner -o ./docs/v2/customer/ --instanceName customer
```
Generate docs employee app files:
```sh
swag init --exclude ./internal/controller/http/v1/services/customer,./internal/controller/http/v1/services/admin,./internal/controller/http/v1/services/partner -o ./docs/v2/employee/ --instanceName employee
```
Generate docs admin app files:
```sh
swag init --exclude ./internal/controller/http/v1/services/customer,./internal/controller/http/v1/services/employee,./internal/controller/http/v1/services/partner -o ./docs/v2/admin/ --instanceName admin
```
Generate docs partner app files:
```sh
swag init --exclude ./internal/controller/http/v1/services/customer,./internal/controller/http/v1/services/employee,./internal/controller/http/v1/services/admin -o ./docs/v2/partner/ --instanceName partner
```

## Overview

### Web framework
[Iris](https://www.iris-go.com/) is an efficient and well-designed, cross-platform, web framework with robust set of features. Build your own high-performance web applications and APIs powered by unlimited potentials and portability.

### Database - ORM
[ent](https://entgo.io/docs/getting-started/) is a simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications with large data-models and sticks with the following principles:

-   Easily model database schema as a graph structure.
-   Define schema as a programmatic Go code.
-   Static typing based on code generation.
-   Database queries and graph traversals are easy to write.
-   Simple to extend and customize using Go templates.


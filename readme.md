# 필독 브랜치 규칙

main - prod, stage - stage , dev - 개발서버

pr 순서 feat => dev => (hotfix/bug)stage => main

1. 최신 dev브랜치에서 feature 만들기
2. dev에 push전 dev pull 받기
3. bug/hotfix 를 제외한 브랜치(ex:feat)로 main/stage에 직접pr금지

# board action 서비스 입니다.

cafe_id 와 cafe_type_id 를 유니크 키로 가집니다.

게시판 별로 행동 권한을 저장합니다.

주기능은

- 게시판 주인의 게시판종류별 권한 설정(읽기,쓰기 ,수정,삭제) - 관리자
- 해당 게시판종류의 권한 조회 - 전체

## entity 구조

```text
board_action{ cafe_id + member_id = unique key 
    id           SERIAL PRIMARY KEY,
    cafe_id      int     ,
    cafe_type_id int     ,
    read_roles   varchar ,
    create_roles varchar ,
    update_roles varchar ,
    update_able  bool ,   // 업데이트가 가능한지 여부 불가일경우 update_roles 가 존재 해도 업데이트 불가입니다. 
    delete_roles varchar ,
    created_at   timestamptz
}
```

### makefile

```shell
# local postgres run (docker-compose)
make local-db
# local postgres migrate init
make local-init
# local postgres apply migrate
make local-migrate
```

# swagger 설정 [출처](https://www.soberkoder.com/swagger-go-api-swaggo/)

## dev 설정

```shell
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
```

## main에

```code
   import (_ "[project명]/docs")
```

```shell
# swagger json 생성   swag init -g [project main path].go
swag init -g cmd/app/main.go
```

## [스웨거 링크](http://localhost:8082/swagger/index.html)
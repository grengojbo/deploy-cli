# Deploy-Cli
Deploy Utility GIT, Docker, Podman


Пароль желательно не передавать через флаг
```shell
./deploy-cli -p <PASSWORD> <COMMAND>
```
Устанавливайте пароль через системные переменнын

 - SECRET_SSH_KEY - Paste your source content of private key (только через системные переменные)
 - SECRET_SSH_PASSPHRASE - Parse PrivateKey With Passphrase (только через системные переменнын)
 - SECRET_SSH_PASSWORD
 - SECRET_SSH_USERNAME

```shell
SECRET_SSH_PASSWORD=zzz23 ./deploy-cli run
```

```shell
./deploy-cli run --host=example.com -w /home/ubuntu/appname --set-env=MY_VAR=223=22,MY=weqwe --dry-run --verbose -c "pwd"
```

Установка системных переменны

```shell
./deploy-cli --set="MY_ENV=aaa,DEPLOY_MY=bbb" --set=MY_ENV3=ccc run
```
Будут переданы переменные окружения
 ```shell
export MY_ENV=aaa; export DEPLOY_MY=bbb; export MY_ENV3=ccc;
```

## Используемые библиотеки

  - https://github.com/robertgzr/porcelain
  - https://github.com/pressly/sup
  - https://github.com/appleboy/easyssh-proxy


# Autoredial system
From [sqlc-buf-pg-template](https://github.com/tk42/sqlc-buf-pg-template)

## To start
 1. Create a Container-Optimized instance on VM. [GCP Container-Optimized OS で docker-compose を使う方法](https://qiita.com/hankehly/items/c7f7c94f50305dce782f)
 2. Added docker-compose alias in .bashrc
 3. Clone this repository
    ```bash
    git clone https://github.com/tk42/auto-redial
    ```
 4. Create DB by
    ```bash
    docker-compose -f docker-compose.dbmate.yml up
    ```
 5. Start the system
    ```bash
    docker-compose up
    ```

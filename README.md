# teleserver

## API约定

1. 客户注册

    请求：`POST /userRegister`

    请求体：

    ```json
    {
        "name": "rammiah", // 用户名
        "pass": "5d41402abc4b2a76b9719d911017c592", // 密码的md5
        "menu": 0 // 套餐
    }
    ```

    响应：

    ```json
    {
        "err": "", // 发生错误时的信息
        "success": true, // 注册成功，注册失败时uid为空字符串
        "uid": "000001" // 用户的id，6位数字
    }
    ```
    
2. 客户登录

    请求：`GET /userLogin?uid=<用户名>&pass=<密码md5摘要>`
    响应：

    ```json
    {
        "success": true, // 是否允许登录
        "name": "rammiah", // 用户名
        "money": 20.0 // 余额
    }
    ```
    
3. 客户查询话费

    请求：`GET /queryConsume?uid=<用户名>&index=<记录下标>&type=<查询类型，对应查询方式>&cluster=<是否分类>`

    响应：

    ```json
    {
        "success": true,
        "err": "",
        "left": 10, // 剩余记录数
        "records": [
            {
                "id": 0, // 流水号
                "year": 2019, // 年
                "month": 5, // 月 
                "day": 1, // 天
                "tm": "20:00:00", // 时间
                "cost": 10.0, // 花费
            },
        ]
    }
    ```

4. 客户查询缴费情况：

    请求：`GET /queryCharge?uid=<用户名>&index=<下标>`

    响应：

    ```json
    {
        "err", "", // 出错时的错误信息
        "success": true,
        "left": 10, // 剩余记录数
        "records": [
            {
                "id": 0, // 流水号
                "cashier_id": "000000", // 收款员ID
                "year": 2019, // 年
                "month": 5, // 月
                "day": 10, // 天
                "tm": "14:10:00", // gin会将数据进行base64转换
                "money": 20.0 // 收款金额
            },
        ]
    }
    ```
    
5. 获取所有可用套餐

    请求：`GET /getMenu`

    响应：

    ```json
    {
        "success": true,
        "menu": [
          	{
                "id": 0,
                "name": "8元套餐",
                "money": 8.0
            }
        ]
    }
    ```

6. 缴费

   请求：`POST /charge`

   请求体：

   ```json
   {
       "cashier_id": "000000", // 收款员id
       "user_id": "000001", // 交款用户id
       "money": 20.0 // 交款金额
   }
   ```

   响应体：

   ```json
   {
       "success": true, // 是否成功
       "err": "" // 如果不成功时返回的错误信息
   }
   ```

7. 收款员登录

    请求：`GET /cashierLogin?uid=<uid>&pass=<pass>`

    响应体：

    ```json
    {
        "success": true, 
        "err": "",
        "name": ""
    }
    ```

8. 客服登录

    请求：`GET /customerServiceLogin?uid=<uid>&pass=<pass>`

    响应体：

    ```json
    {
        "success": true,
        "err": "",
        "name": ""
    }
    ```

    

    大体类似客户登录

9. 客服的终端发送服务记录

    请求：`POST /service`

    请求体：

    ```json
    {
        "customer_service_id": "000000", // 客服的编号
        "user_id": "000001", // 用户的编号
    }
    ```

    响应体：

    ```json
    {
        "success": true, // 记录是否添加成功
        "err": "" // 失败时的错误信息
    }
    ```

10. 用户ID有效性检测

    请求：`GET /validUserId?uid=<uid>`

    响应体：

    ```json
    {
        "err": "",
        "menu": 0,
        "menu_name": "30元套餐",
        "money": 0.0,
        "name": "",
        "success": true,
        "valid": true
    }
    ```

11. 套餐修改

     请求：`POST /changeMenu`

     请求体：

     ```json
     {
         "uid": "000000",
         "new_menu": 0
     }
     ```

     响应体：

     ```json
     {
         "err": "",
         "success": true
     }
     ```

12. 消费接口

     请求：`POST /consume`

     请求体：

     ```json
     {
         "uid": "000000",
         "cost": 10.0
     }
     ```

     响应体：

     ```json
     {
         "err": "",
         "success": true
     }
     ```

13. 客服数据统计

     请求：`GET /serviceStatistics?year=<year>&month=<month>`

     响应体：

     ```json
     {
         "err": "", // 请求失败时的错误信息
         "records": [
             {
                 "ser_id": "000000",
                 "cnt": 3,
                 "name": "沈翠玲"
             },
             {
                 "ser_id": "000001",
                 "cnt": 0,
                 "name": "田元亮"
             },
             {
                 "ser_id": "000002",
                 "cnt": 20,
                 "name": "徐瑜"
             }
         ],
         "success": true // 请求是否成功
     }
     ```

14. 管理员登陆

     请求：`GET /adminLogin?uid=<uid>&pass=<pass>`

     响应体：

     ```json
     {
         "err": "",
         "name": "rammiah",
         "success": true
     }
     ```

15. 密码重置

     请求：`POST resetPassword`

     请求体：

     ```json
     {
         "uid": "000000",
         "type": 0, // 表示重置人员类型，用户，客服，收款员
         "pass": "5c5ed39577288f31880701799bdd7cde"
     }
     ```

     

# interpreterInGo

使用Go语言实现一个简单的解释器



# 使用方法

    go run main.go

# 执行单测
    
        go test ./lexer

# 语言特性
定义了一个新的语言，该语言的语法如下：

    ```
    let x = 5;
    let y = 10;
    let add = fn(x, y) {
        x + y;
    };
    let result = add(x, y);

    ```
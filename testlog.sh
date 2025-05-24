#!/bin/bash

# 文件名
file="logs/book.log"

# 初始化文件内容为空
#> "$file"

# 循环条件初始化
continue_loop=true

# 模拟写入数据到文件
while $continue_loop; do
    curl http://localhost:8080/api/books/3
    
    # 获取文件大小（单位：字节）
    file_size=$(stat -f%z "$file")
    
    # 将文件大小转换为兆字节（MB）
    file_size_mb=$(echo "scale=2; $file_size / 1024 / 1024" | bc)
    
    # 检查文件大小是否超过1.1M
    if (( $(echo "$file_size_mb > 1.1" | bc -l) )); then
        echo "File size exceeds 1.1 MB. Stopping the loop."
        continue_loop=false
    fi
    
    # 延迟一段时间以观察文件增长
    sleep 1
done

echo "Script finished."

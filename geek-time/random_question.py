import re
import random

def extract_random_subheadings(file_path, count=10):
    try:
        with open(file_path, 'r', encoding='utf-8') as file:
            content = file.read()
        
        # 使用正则表达式匹配所有的 ## 二级标题
        subheadings = re.findall(r'^##\s+(.+)', content, re.MULTILINE)
        
        if not subheadings:
            print("No second-level headings (##) found in the file.")
            return None
        
        # 随机选择最多 `count` 个二级标题
        selected_subheadings = random.sample(subheadings, min(count, len(subheadings)))
        
        print(f"Randomly selected subheadings ({len(selected_subheadings)}):")
        for subheading in selected_subheadings:
            print(f"- {subheading}")
        
        return selected_subheadings
    except FileNotFoundError:
        print(f"File not found: {file_path}")
    except Exception as e:
        print(f"An error occurred: {e}")

# 使用示例
markdown_file_path = 'mysql.md'  # 替换为你的 Markdown 文件路径
extract_random_subheadings(markdown_file_path, count=10)

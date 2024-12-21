import os
import re
import random

def extract_questions_and_answers(base_dirs, output_file, num_questions=3):
    """
    从指定目录中的 Markdown 文件中抽取以 ## 开头的问题及答案（问题到下一个 ## 标题之间的内容），并保存到一个新的文件中。
    尽量让问题分布在多个文件中。

    :param base_dirs: 要搜索的目录列表
    :param output_file: 输出的 Markdown 文件
    :param num_questions: 抽取的问题数量
    """
    questions_and_answers = []

    # 遍历目录并提取问题和答案
    for base_dir in base_dirs:
        for root, _, files in os.walk(base_dir):
            for file in files:
                if file.endswith(".md"):
                    filepath = os.path.join(root, file)
                    with open(filepath, 'r', encoding='utf-8') as f:
                        lines = f.readlines()

                    question_pattern = re.compile(r'^##\s+(.*)')
                    question_start = None
                    for i, line in enumerate(lines):
                        match = question_pattern.match(line)
                        if match:
                            # 保存之前的问题及答案
                            if question_start is not None:
                                questions_and_answers.append({
                                    'file': filepath,
                                    'question': lines[question_start].strip(),
                                    'answer': "".join(lines[question_start + 1:i]).strip(),
                                    'line_number': question_start + 1,
                                })
                            question_start = i
                    # 保存最后一个问题及答案
                    if question_start is not None:
                        questions_and_answers.append({
                            'file': filepath,
                            'question': lines[question_start].strip(),
                            'answer': "".join(lines[question_start + 1:]).strip(),
                            'line_number': question_start + 1,
                        })

    # 对问题按文件进行分组
    questions_by_file = {}
    for qa in questions_and_answers:
        questions_by_file.setdefault(qa['file'], []).append(qa)

    # 从每个文件随机选一个问题，直到达到 num_questions 或用尽问题
    selected_questions = []
    while len(selected_questions) < num_questions and questions_by_file:
        for file, questions in list(questions_by_file.items()):
            if questions:
                selected_questions.append(questions.pop(random.randint(0, len(questions) - 1)))
            if not questions:  # 如果文件的问题用尽，移除文件
                del questions_by_file[file]

        # 如果问题已达到目标数量，提前退出
        if len(selected_questions) >= num_questions:
            break

    # 写入输出文件
    with open(output_file, 'w', encoding='utf-8') as f:
        for qa in selected_questions:
            f.write(f"### 问题: {qa['question'][3:]}\n")  # 保留问题标题并去掉 "## " 前缀
            f.write(f"*来源*: `{qa['file']}:{qa['line_number']}`\n\n")
            f.write(f"**答案**:\n{qa['answer']}\n\n")
            f.write("---\n\n")

    print(f"问题和答案已保存到 {output_file}")

# 使用示例
base_dirs = ['blog', 'geek-time2']
output_file = 'extracted_questions_with_content.md'
extract_questions_and_answers(base_dirs, output_file)

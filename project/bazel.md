## Bazel优点
- 多cache层，编译非常快，增量构建，并发量大
- 高准确性，actions都是在单独的沙盒里，相同的源码总是得到相同的二进制
- 可扩展
- 多语言支持，C++，go, java, python, rust
- 多平台，交叉编译能力

## Bazel不足
- 编译过程中input files变化会导致上传无效的结果到remote cache
- 在Docker 容器中执行建构作业时，内存中的增量状态会遗失：
    即使在单一Docker 容器中执行，Bazel 也会使用sever/clirnt架构。在server端，Bazel 会维护状态，加快建构作业。在Docker 容器(例如CI) 中执行建构作业时，内存中的状态会遗失，因此Bazel 必须先重建内存内的状态，才能使用remote cache。
    - 使用PVC runner来优化
    - 同时也是项目需求，谈一下背景
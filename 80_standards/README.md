# 规范

一个应用基本都是多人协作开发的，但不同开发者的开发习惯、方式都不同。如果没有一个统一的规范，就会造成非常多的问题，如：

- 代码风格不一：代码仓库中有多种代码风格，读/改他人的代码都是一件痛苦的事情，整个代码库也会看起来很乱。
- 目录杂乱无章：相同的功能被放在不同的目录，或一个目录根本不知道它要完成什么功能，新开发的代码也不知道放在哪个目录或文件。这些都会严重降低代码的可读性和可维护性。
- 接口不统一：对外提供的 API 接口不统一，如，修改用户接口为 `/v1/users/admin`，但是修改密钥接口为 `/v1/secret?name=secret0`，难以理解和记忆。
- 错误码不规范：错误码格式不统一，导致难以辨别错误类型；同类错误拥有不同错误码，增加理解难度。

因此，我们需要一个好的规范来约束开发者，以确保大家开发的是“**一个应用**”。一个好的规范不仅可以提高软件质量，也可以提高软件的开发效率、降低维护成本，甚至减少 Bug 数，还可以使开发者的体验如行云流水一般顺畅。所以，在编码之前，有必要花一些时间和团队成员一起讨论并制定规范。

一个 Go 应用会涉及很多方面的规范，同类规范也会因为团队差异而有所不同。所以，本章只介绍一些开发中常用的规范，包括日志规范、错误规范、代码规范和 Git 工作流规范。

## 日志

- [日志](10_log/README.md)

## 错误处理

- [错误处理](20_error/README.md)

## 代码规范

- [代码规范](30_code/README.md)

## 工作流程

- [工作流程](60_workflow/README.md)

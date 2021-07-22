# Learn Scala

I'm interested in scala because of it's Python-like syntax and it has static typings.

What turns me down was when I saw "run on top of JVM" lol.

Here's what you'll need:
* OpenJDK (I have v11)
* SBT (download here https://www.scala-sbt.org/download.html)

To run it, simply:
```bash
# Let's try the basic directory
$ cd basic
$ sbt run
# [info] welcome to sbt 1.5.5 (Ubuntu Java 11.0.11)
# [info] loading project definition from /home/reinaldyrfl/repository/learn/scala/basic/project
# [info] loading settings for project root from build.sbt ...
# [info] set current project to scala3-simple (in build file:/home/reinaldyrfl/repository/learn/scala/basic/)
# [info] compiling 3 Scala sources to /home/reinaldyrfl/repository/learn/scala/basic/target/scala-3.0.1/classes ...
# [info] running basic.hello 
# Hello world!
# I was compiled by Scala 3. :)
# 10
# 10.0
# 0
# 1
# 2
# 3
# 4
# ...
# [success] Total time: 5 s, completed Jul 23, 2021, 12:44:58 AM
```
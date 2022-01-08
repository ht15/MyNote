valgrind  YOUR_ELF --tool=memcheck --log-file=valgrind_${now}.log --leak-check=full --show-leak-kinds=definite --time-stamp=yes -v 


#include "zend.h"
void php_set_ini(char *ini);
void php_startup();
char * php_exec_file(char *filename);
char * php_exec_error();
void php_shutdown(void);

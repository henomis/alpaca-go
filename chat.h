#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>


int alpaca(
    int help, 
    int interactive,
    int interactive_start,
    char *reverse_prompt,
    int color,
    int seed,
    int threads,
    char *prompt
);


#ifdef __cplusplus
}
#endif
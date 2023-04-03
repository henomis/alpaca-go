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
    char *prompt,
    char *prompt_file,
    int n_predict,
    int top_k,
    double top_p,
    int repeat_last_n,
    double repeat_penalty,
    int prompt_size,
    double temperature,
    int batch_size,
    char *model_file
);


#ifdef __cplusplus
}
#endif
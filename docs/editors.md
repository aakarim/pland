# Editor support
## Vi(m)
### Background updating
Pland updates your .plan file in the background. If you have an open buffer it won't update by default. You'll either need to update your config file with `set autoread` or run `:checktime` so you can see if it has changed. See this Stackoverflow https://unix.stackexchange.com/questions/149209/refresh-changed-content-of-file-opened-in-vim

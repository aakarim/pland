# How to edit your .plan
A .plan file is a UTF-8 text file with a few special tokens that allow it to be understood by our servers. You can consider it like a .txt and can open it with any text editor.

# Editor support
## Vi(m)
### Background updating
Pland updates your .plan file in the background. If you have an open buffer it won't update by default. You'll either need to update your config file with `set autoread` or run `:checktime` so you can see if it has changed. See this Stackoverflow https://unix.stackexchange.com/questions/149209/refresh-changed-content-of-file-opened-in-vim

" 行末の空白をハイライト
augroup HighlightTrailingSpaces
  autocmd!
  autocmd VimEnter,WinEnter,ColorScheme * highlight TrailingSpaces term=underline guibg=Red ctermbg=1
  autocmd VimEnter,WinEnter * match TrailingSpaces /\s\+$/
augroup END

set tabstop=4
set shiftwidth=4
set expandtab
au BufNewFile,BufRead *.yaml,*.yml   set expandtab tabstop=2 shiftwidth=2
au BufNewFile,BufRead Makefile,*mk   set noexpandtab tabstop=2 shiftwidth=2
au BufNewFile,BufRead *.go           set noexpandtab tabstop=4 shiftwidth=4

" 行番号とカレント行
set cursorline
"hi CursorLine   term=reverse cterm=none ctermbg=240
hi CursorLine   term=reverse cterm=none ctermbg=none
hi CursorLineNr term=bold    cterm=none ctermfg=228 ctermbg=none
set number
hi LineNr ctermfg=59

" 121桁目から背景を変える
let &colorcolumn=join([51,81,91,121],",")
hi ColorColumn guibg=White

" 新しいウィンドウを下に開く
set splitbelow
" 新しいウィンドウを右に開く
set splitright

"===============================================================================
"$ mkdir -p ~/.vim/colors/
"$ curl --location --output ~/.vim/colors/gruvbox.vim https://raw.githubusercontent.com/morhetz/gruvbox/master/colors/gruvbox.vim
"===============================================================================
syntax on
syntax enable
set background=light
colorscheme desert
if filereadable(expand('~/.vim/colors/gruvbox.vim'))
  colorscheme gruvbox
endif

set fileformats=unix,dos,mac

"===============================================================================
" Install方法
"$ mkdir -p ~/.vim/pack/plugins/start/
"$ git clone git://github.com/nathanaelkane/vim-indent-guides.git ~/.vim/pack/plugins/start/vim-indent-guides
"===============================================================================
if isdirectory(expand('~/.vim/pack/plugins/start/vim-indent-guides'))
  let g:indent_guides_enable_on_vim_startup = 1
  let g:indent_guides_auto_colors=0
  let g:indent_guides_start_level=1
  let g:indent_guides_guide_size=1
  hi IndentGuidesOdd  ctermbg=239
  hi IndentGuidesEven ctermbg=241
endif

"===============================================================================
" Install方法(vim-lspとvim-lsp-settingsの2つ)
" 設定必要なし(*.goを開くとgoが入っていれば提案してくれる)
"$ mkdir -p ~/.vim/pack/plugins/start/
"$ git clone https://github.com/prabirshrestha/vim-lsp.git ~/.vim/pack/plugins/start/vim-lsp
"$ git clone https://github.com/mattn/vim-lsp-settings.git ~/.vim/pack/plugins/start/vim-lsp-settings
"===============================================================================

"===============================================================================
" Install方法
"$ mkdir -p ~/.vim/pack/plugins/start/
"$ git clone https://github.com/mattn/vim-goimports.git ~/.vim/pack/plugins/start/vim-goimports
"===============================================================================
if isdirectory(expand('~/.vim/pack/plugins/start/vim-goimports'))
  let g:goimports = 1
  let g:goimports_simplify = 1
endif


"===============================================================================
" vim-lspのGo
" https://github.com/prabirshrestha/vim-lsp/wiki/Servers-Go
"===============================================================================
if executable('gopls')
  au User lsp_setup call lsp#register_server({
    \ 'name': 'gopls',
    \ 'cmd': {server_info->['gopls']},
    \ 'whitelist': ['go'],
    \ })
  autocmd BufWritePre *.go LspDocumentFormatSync
endif

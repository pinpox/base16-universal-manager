# Fish completions for base16-universal-manager

# Never complete with filenames
complete -f -c base16-universal-manager -d "A command line tool to install base16 templates and set themes globally"

# Complete available themes
set -l themes_list (sed -n 's/^\(.*\).yaml:.*/\1/p' ~/.cache/base16-universal-manager/schemeslist.yaml)
complete -f -c base16-universal-manager -n "__fish_seen_subcommand_from --scheme; and test (count (commandline -opc)) -eq 2" -a "$themes_list"

# Description for each command
set -l base16_universal_manager_commands (base16-universal-manager --help-long | sed -En 's/^  (--(\w-?)+).*/\1/p')
set -l base16_universal_manager_commands_help (base16-universal-manager --help-long | sed 1,3d | sed -En 's/.{21}([A-Z].*)/\1/p')
set -l i 1
for command in $base16_universal_manager_commands
    complete -f -c base16-universal-manager -n "not __fish_seen_subcommand_from $base16_universal_manager_commands" -a "$command" -d "$base16_universal_manager_commands_help[$i]"
    set -l i (math $i + 1)
end

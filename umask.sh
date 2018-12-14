for user in $(awk -F: '{print $1}' /etc/passwd);
do
    printf "%-10s" "$user" ; su -c 'umask' -l $user 2>/dev/null
done

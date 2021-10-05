awk '{if(NF==1)print $1; else print $4, $3, $2, $1}'

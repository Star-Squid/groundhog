#!/bin/bash
# find . -type f -name '*.sh' | cut -d "." -f2 | cut -d "/" -f2 | sort -r -t '\0' -n 
ls -R | grep '.*[.]sh' | cut -d "." -f1 | sort -r -t '\0' -n
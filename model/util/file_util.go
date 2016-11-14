package util

import (
    "bufio"
)

func FileReadline(in *bufio.Reader) (string) {
    
    var cur_line = make([]rune, 100);

    for {
        temp_rune,_,err := in.ReadRune();

        if err != nil || temp_rune == '\n' || temp_rune == '\r' {
            break;
        }

        cur_line = append(cur_line, temp_rune);
    }

    line := string(cur_line);

    return line;
}

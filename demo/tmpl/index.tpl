axxxxxxxxxxxxx

{{.user_list}}
{{index .test_data "name" }}
{{ccTest "nice"}}

<html>
    <head> </head>
    <body>
        <table>
            <tr>
                <th>id</th>
                <th>名字</th>
                <th>年龄</th>
                <th>生日</th>
            </tr>

            {{ range $i, $user := .user_list }}
                {{if isRightType $user}}
                <tr>
                    <td>{{$i}}</td>
                    <td>{{index $user "name"}}</td>
                    <td>{{.test_data}}</td>
                </tr>
                {{end}}
            {{end}}

        </table>
    </body>
</html>

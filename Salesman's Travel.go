package kata

import (
    "fmt"
    "strings"
)

func Travel(r, zipcode string) string {
    rList := strings.Split(r, ",")
    addressList := make([]string, 0)
    houseNumberList := make([]string, 0)
    if zipcode != "" {

        for _, s := range rList {
            if strings.HasSuffix(s, zipcode) {
                // 测试集里包括了换行……
                address := strings.Trim(strings.TrimSuffix(s, zipcode), " \n\r\t")
                list := strings.Split(address, " ")
                houseNumber := list[0]
                address = strings.Join(list[1:], " ")
                addressList = append(addressList, address)
                houseNumberList = append(houseNumberList, houseNumber)
            }
        }
    }
    addresses := strings.Join(addressList, ",")
    houseNumbers := strings.Join(houseNumberList, ",")
    return fmt.Sprintf("%s:%s/%s", zipcode, addresses, houseNumbers)
}

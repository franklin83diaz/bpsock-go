package bpsock

import (
	"bpsock-go/utils"
	"fmt"
)

func (bpsock *BpSock) received() {

	start := 0
	end := 22
	buffer := make([]byte, bpsock.dmtu+22)
	var idChan int
	var tagName string
	var sizeData int

	for {

		// Read data
		bytesRead, err := bpsock.socket.Read(buffer[start:end])

		if err != nil {
			//if the error is EOF, then the socket is closed
			// no need to print the error
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading data: ", err)
			break
		}

		sizeUnit := bytesRead

		if start == 0 {
			b := buffer[:bytesRead]
			//id
			idBytes := b[0:2]
			idChan = int(idBytes[0])<<8 | int(idBytes[1])

			//tag
			tagBytes := b[2:18]
			tagName = utils.BytesToStringTrimNull(tagBytes)

			//size data
			sizeDataBytes := b[18:22]
			sizeData = int(sizeDataBytes[0])<<24 | int(sizeDataBytes[1])<<16 | int(sizeDataBytes[2])<<8 | int(sizeDataBytes[3])

			//sizeUnit is the size of the data plus the header
			sizeUnit = sizeData + 22
		}

		//if the size of the data is greater than the bytes read
		if sizeUnit > bytesRead {
			start = bytesRead
			end = sizeData + 22
			continue
		}

		//reset start and end
		start = 0
		end = 22

		//if is end channel

		//data
		data := make([]byte, len(buffer[22:sizeData+22]))
		copy(data, buffer[22:sizeData+22])

		tagNameOrig := ""
		//check if is request
		if tagName[0] == '1' {
			tagNameOrig = tagName[1:]
			tagName = tagName[8:]
		}
		if tagName[0] == '2' {
			tagName = tagName[1:]
		}

		//get the handlers
		listHandlers := bpsock.handlers

		//check if the tag is in the list of handlers
		for i := 0; i < len(listHandlers); i++ {

			//if the tag is in the list of handlers
			if listHandlers[i].Tag().Name() == tagName {

				//if sizeData is 0, then it is an end channel
				if sizeData == 0 {
					action := listHandlers[i].ActionFunc()

					//if is request
					if tagNameOrig != "" {

						go func() {
							//remove data from the handler after the action is executed
							defer listHandlers[i].RemoveData(idChan)
							action(listHandlers[i], tagNameOrig, idChan)
						}()

					} else { //else is a async

						go func() {
							//remove data from the handler after the action is executed
							defer listHandlers[i].RemoveData(idChan)
							action(listHandlers[i], tagName, idChan)
						}()

					}

					break
				}

				//add data to the handler
				listHandlers[i].AddData(idChan, data)

				//just one handler per tag, no continue to the next handler
				break

			}
		}

	}
}

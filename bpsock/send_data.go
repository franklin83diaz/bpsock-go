package bpsock

import (
	//lint:ignore ST1001 import tags
	. "bpsock-go/tags"
	"bpsock-go/utils"
	"net"
	"strconv"
)

func SendData(data []byte, tag Tag16, id_chan int, socket net.Conn, dmtu int, typeSend ...int) error {
	/// 22 bytes for the header of Unit (id, tag, sizeData)

	// id_chan 2 bytes
	bytesId := make([]byte, 2)
	bytesId[0] = byte(id_chan >> 8)
	bytesId[1] = byte(id_chan)

	// tag 16 bytes
	tagName := tag.Name()
	//add typeSend to tag
	if len(typeSend) > 0 {
		tagName = strconv.Itoa(typeSend[0]) + tagName
	}

	bytesTag := make([]byte, 16)
	copy(bytesTag, []byte(tagName))

	// sizeData 4 bytes
	sizeData := len(data)
	bytesSizeData := make([]byte, 4)
	bytesSizeData[0] = byte(sizeData >> 24)
	bytesSizeData[1] = byte(sizeData >> 16)
	bytesSizeData[2] = byte(sizeData >> 8)
	bytesSizeData[3] = byte(sizeData)

	/// if is bigger
	if sizeData > dmtu {

		splitData := utils.SplitData(data, dmtu)

		for _, chunk := range splitData {

			lenData := len(chunk)
			bytesSizeData[0] = byte(lenData >> 24)
			bytesSizeData[1] = byte(lenData >> 16)
			bytesSizeData[2] = byte(lenData >> 8)
			bytesSizeData[3] = byte(lenData)

			unit := append(bytesId, bytesTag...)
			unit = append(unit, bytesSizeData...)
			unit = append(unit, chunk...)

			// send the unit
			_, err := socket.Write(unit)
			if err != nil {
				return err
			}
		}

		//Send end channel
		//send id_chan add tag
		endChannel := append(bytesId, bytesTag...)
		//send sizeData 0
		endChannel = append(endChannel, []byte{0, 0, 0, 0}...)
		_, err := socket.Write(endChannel)
		if err != nil {
			return err
		}

	} else {

		// create the unit
		unit := append(bytesId, bytesTag...)
		unit = append(unit, bytesSizeData...)
		unit = append(unit, data...)

		// send the unit
		_, err := socket.Write(unit)
		if err != nil {
			return err
		}

		//if is cancel no send end channel
		// tag end channel start with runner "3"
		if tagName[0] == 3 {
			return nil
		}

		// send the end channel
		// send id_chan + tag end channel + sizeData 0
		//send id_chan add tag
		endChannel := append(bytesId, bytesTag...)
		//send sizeData 0
		endChannel = append(endChannel, []byte{0, 0, 0, 0}...)
		_, err = socket.Write(endChannel)
		if err != nil {
			return err
		}

	}

	return nil

}

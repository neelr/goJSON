import React from "react"
import { Flex, Heading, Button } from "rebass"
import { Input } from "@rebass/forms"
import ls from "local-storage"

export default class extends React.Component {
    tokenInput = null
    connect() {
        if (this.tokenInput.value != "") {
            let tokenArray = []
            if (ls("tokenList")) {
                tokenArray = JSON.parse(ls("tokenList"))
            }
            tokenArray.unshift(this.tokenInput.value.replace(" ", ""))
            ls("tokenList", JSON.stringify(tokenArray))
            window.location.href = "/viewer"
        }
    }
    render() {
        return (
            <Flex flexDirection="column" mx="auto">
                <Flex mt="100px" flexDirection="column">
                    <Heading fontSize={[2, 3, 4]} py="20px">Connect to New Database!</Heading>
                    <Input mb="20px" ref={x => this.tokenInput = x} placeholder="Token" onKeyPress={e => {
                        if (e.key == "Enter") {
                            this.connect()
                        }
                    }} />
                    <Button sx={{
                        bg: "highlight",
                        m: "20px",
                        my: "auto",
                        ":hover": {
                            bg: "secondary",
                            cursor: "pointer"
                        }
                    }} onClick={x => this.connect()}>Connect 🔌</Button>
                </Flex>
            </Flex>
        )
    }
}
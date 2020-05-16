import React from "react"
import { Flex, Heading, Text } from "rebass"
import ls from "local-storage"

export default class extends React.Component {
    render() {
        return (
            <Flex sx={{
                width: [0, "35vw", "35vw", "25vw"],
                bg: "secondary",
                height: "100%",
                transition: "all 0.5s",
                flexDirection: "column",
                overflow: "scroll"
            }}>
                <Heading fontSize={[1, 2, 3]} color="white" mx="auto" my="20px" mt="70px">Connections</Heading>
                {
                    ls("tokenList") ? JSON.parse(ls("tokenList")).map(v => (
                        <Flex sx={{
                            ":hover": {
                                bg: "highlight",
                                cursor: "pointer"
                            },
                            color: "white",
                            p: "10px"
                        }} onClick={() => {
                            let tokenArray = []
                            if (ls("tokenList")) {
                                tokenArray = JSON.parse(ls("tokenList"))
                            }
                            tokenArray.unshift(v)
                            ls("tokenList", JSON.stringify(tokenArray))
                            window.location.href = "/viewer"
                        }}>
                            <Text mx="auto">{v.length > 15 ? `${v.substring(0, 14)}...` : v}</Text>
                        </Flex>
                    )) : null
                }
            </Flex>
        )
    }
}
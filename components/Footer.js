import React from "react"
import { Flex, Text, Link } from "rebass"
const A = (props) => (
    <Link sx={{
        color: "blue.5",
        textDecorationStyle: "wavy"
    }} {...props} />
)
export default () => (
    <Flex mt="auto" height="100px">
        <Text fontWeight={700} m="auto">Made with ☕ by <A href="https://github.com/neelr">@neelr</A> | <A href="https://github.com/neelr/goJSON/tree/dash">Source</A></Text>
    </Flex>
)
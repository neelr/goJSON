import { Flex, Text } from "rebass"
import Link from "next/link"

const Item = ({ sx, children, ...props }) => (
    <Flex sx={props.hover ? {
        fontWeight: 700,
        color: "white",
        fontSize: 2,
        px: "20px",
        ":hover": {
            bg: "secondary",
            cursor: "pointer"
        },
        a: {
            textDecoration: "none",
            color: "white",
        },
        ...sx
    } : {
            fontWeight: 700,
            color: "white",
            fontSize: 2,
            px: "20px",
            a: {
                textDecoration: "none",
                color: "white",
            },
            ...sx
        }} {...props}>
        <Text my="auto">{children}</Text>
    </Flex>
)

export default (props) => (
    <Flex sx={{
        bg: "primary",
        height: "50px",
        position: "fixed",
        top: 0,
        width: "100vw"
    }} {...props}>
        <Item>
            goJSON
        </Item>
        <Item ml="auto" fontSize={1} hover>
            <Link href="/">
                + New Connection
            </Link>
        </Item>
    </Flex>
)
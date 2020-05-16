import React from "react";
import { Flex, Heading, Button, Text } from "rebass";
import ls from "local-storage";
var ReactJson = null;

export default class extends React.Component {
    state = {
        json: {},
        name: null,
        status: null,
    };
    render() {
        return (
            <Flex flexDirection="column" mx="auto">
                <Flex mt="100px" flexDirection="column">
                    {this.state.status == "Done!" ? <Text color="green.3">{this.state.status}</Text> : <Text color="red.5">{this.state.status}</Text>
                    }
                    <Heading fontSize={[1, 2, 3]} py="20px">
                        Connected to {this.state.name}!
					</Heading>
                    {ReactJson ? (
                        <ReactJson
                            onEdit={(json) =>
                                this.setState({ json: json.updated_src })
                            }
                            onAdd={(json) =>
                                this.setState({ json: json.updated_src })
                            }
                            theme="summerfruit:inverted"
                            iconStyle="circle"
                            src={this.state.json}
                            onDelete={(json) =>
                                this.setState({ json: json.updated_src })
                            }
                        />
                    ) : null}
                    <Button
                        sx={{
                            color: "white",
                            bg: "green.4",
                            mt: "20px",
                            m: "auto",
                            ":hover": {
                                bg: "secondary",
                                cursor: "pointer",
                            },
                        }}
                        onClick={() => {
                            console.log(this.state.json)
                            this.setState({
                                status: "",
                            })
                            fetch(
                                `https://db.neelr.dev/api/${this.state.name}`,
                                {
                                    method: "POST",
                                    body: JSON.stringify(this.state.json),
                                    headers: {
                                        "Content-Type": "application/json",
                                    },
                                }
                            )
                                .then((d) => this.setState({ status: "Done!" }))
                                .catch((d) =>
                                    this.setState({
                                        status: "Error! Try Again in a bit.",
                                    })
                                );
                        }}
                    >
                        Save
					</Button>
                </Flex>
            </Flex>
        );
    }
    componentDidMount() {
        require.ensure(["react-json-view"], function () {
            try {
                ReactJson = require("react-json-view").default;
            } catch (err) {
                console.log("react-json-view:", err);
            }
        });
        ls("tokenList") ? null : window.location.href = "/"
        fetch(`https://db.neelr.dev/api/${JSON.parse(ls("tokenList"))[0]}`)
            .then((body) => body.json())
            .then((r) => {
                this.setState({
                    json: r,
                    name: JSON.parse(ls("tokenList"))[0],
                });
            })
            .catch((d) => {
                this.setState({ name: JSON.parse(ls("tokenList"))[0] });
            });
    }
}

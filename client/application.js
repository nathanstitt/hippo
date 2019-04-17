import React from 'react';
import whenDomReady from 'when-dom-ready';
import { render } from 'react-dom';
import ApolloClient from 'apollo-boost';

import { ApolloProvider } from 'react-apollo';
import { gql } from 'apollo-boost';
import { Query } from 'react-apollo';
import {
    App, Split, Sidebar, Header, TItle, Box, Menu, Anchor,
} from 'grommet'


const Application = () => (
    <App centered={false}>
        <Split separator={true} flex='right' fixed={true} colorIndex='brand'>
            <Sidebar colorIndex='brand' responsive={true}>
                <Header pad='small' justify='between'>
                    <Title> brand </Title>
                </Header>
                <Box flex='grow' justify='start'>
                    <Menu primary={true}>
                        <Anchor key="home" path="/app" label="Home" />
                    </Menu>
                </Box>
            </Sidebar>
            <Box>
                {this.props.children}
            </Box>
        </Split>
    </App>
)


const GET_USER = gql`
query {
    users(where: {name: {_eq: "nathan"}}) {
        id
        name
        created_at
    }
}
`

// const App = () => (
// <Query query={GET_USER}>
//     {({ loading, error, data }) => {
//          console.log(data)
//          if (loading) return <div>Loading...</div>;
//          if (error) return <div>Error...</div>;
//
//          return data.users.map(u => <div key={u.id}>{u.id}: {u.name}</div>)
//     }}
// </Query>
// )


const ApolloApp = (AppComponent, JWT) => {
    // Pass your GraphQL endpoint to uri
    const client = new ApolloClient({
        uri: '/v1alpha1/graphql',
        headers: {
            Authorization: `Bearer ${JWT}`,
        },
    });
    return (
        <ApolloProvider client={client}>
            <AppComponent />
        </ApolloProvider>
    );
};

whenDomReady(() => {
    const bootstrapData = JSON.parse(
        document.getElementById('bootstrapData')
            .getAttribute('content')
    );
    console.log(bootstrapData)
    const JWT = bootstrapData.JWT;
    render(ApolloApp(Application, JWT), document.getElementById('root'));
});

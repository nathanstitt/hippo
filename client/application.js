import React from 'react';
import whenDomReady from 'when-dom-ready';
import { render } from 'react-dom';
import ApolloClient from 'apollo-boost';

import { ApolloProvider } from 'react-apollo';
import { gql } from 'apollo-boost';
import { Query } from 'react-apollo';


const GET_USER = gql`
    query {
        users(where: {name: {_eq: "nathan"}}) {
            id
            name
            created_at
        }
    }
`

const App = () => (
    <Query query={GET_USER}>
        {({ loading, error, data }) => {
             console.log(data)
             if (loading) return <div>Loading...</div>;
             if (error) return <div>Error...</div>;

             return data.users.map(u => <div key={u.id}>{u.id}: {u.name}</div>)
        }}
    </Query>
)


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
    render(ApolloApp(App, JWT), document.getElementById('root'));
});

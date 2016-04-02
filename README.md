# valente - websocket for golang webapp

valente is an experiment with Golang webapp using WebSockets. There are similar solutions in another languages/platform:

 * The Wt WebFramework has an experimental features with all comunications between server and browser using websockets. 
 * N2O and Nitrogen are frameworks with that feature coded in Erlang. 

The use of asynchronism is mandatory for this solutions. Traditional threads aren't viable, due to high memory consumed for each connection. 
The Wt Framework is using Boost::asio to handle connections. In Go and Erlang, the languages features should make that scheme simple and scalable. 
I don't know :)

The valente is based on Nitrogen ideas.

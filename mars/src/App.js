import {useState} from "react";

const API_URL = "http://localhost:8000";

function App() {
    // output state
    const [output, setOutput] = useState('');

    function handlePing() {
        setOutput('Hello World!');
    }

    return (
        <div className="App">
            <header className="App-header">
                Crucible
            </header>

            <div className="App-body">
                <button onClick={handlePing}>Ping</button>

                <div className="output-console">
                    {output}
                </div>
            </div>
        </div>
    );
}

export default App;

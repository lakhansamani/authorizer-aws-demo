import { BrowserRouter } from 'react-router-dom';
import { AuthorizerProvider } from '@authorizerdev/authorizer-react';
import App from './App';

export default function Root() {
	return (
		<BrowserRouter>
			<AuthorizerProvider
				config={{
					authorizerURL: 'https://auth.aws-demo.authorizer.dev',
					redirectURL: window.location.origin,
					clientID: '7de5745b-bea2-4696-a825-bdb829836cce',
				}}
			>
				<App />
			</AuthorizerProvider>
		</BrowserRouter>
	);
}

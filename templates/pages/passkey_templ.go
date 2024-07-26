// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/pynezz/pynezz_com/templates/layout"

const formDiv = "flex flex-col justify-content space-between align-items-center bg-bg self-center text-txt p-4 pr-9 pl-9 max-w-50 border border-surface rounded-4 m-0-auto flex-shrink-1 mt-8"

func Passkey() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{"section", formDiv}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/passkey.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 = []any{"h1", layout.Title}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var4...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var4).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/passkey.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Passkey</h1><div class=\"container flex justify-center items-center vh-100\"><div class=\"bg-base p-5 rounded-lg border-crust border-2 w-50 h-max\"><h1 class=\"mb-4 text-center\">🔑 Passkey</h1><div class=\"text-left text-wrap\" id=\"message\"><span id=\"elem-error\"></span> <span id=\"elem-success\"></span></div><div class=\"mb-3 flex flex-col w-full gap-2\"><input type=\"text\" class=\"focus:outline-none bg-crust rounded-md p-2\" id=\"username\" placeholder=\"username\"> <input type=\"text\" class=\"focus:outline-none bg-crust rounded-md p-2\" id=\"displayname\" placeholder=\"display name\"></div><div class=\"flex flex-col w-full\"><div class=\"rounded border-mantle gap-2\"><button class=\"p-2 py-4 bg-surface2 w-full h-fit shadow-md focus:outline-none after:bg-mantle text-text active:bg-surface0 rounded-sm hover:bg-sky hover:text-surface2  \" id=\"registerButton\">Register</button> <button class=\"p-2 bg-surface2 w-full h-fit shadow-sm focus:outline-none aria-pressed:bg-mauve text-text active:bg-surface0 rounded-sm hover:bg-sky hover:text-surface2\" id=\"loginButton\">Login</button></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = pkScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func pkScript() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_pkScript_da18`,
		Function: `function __templ_pkScript_da18(){const { startRegistration } = SimpleWebAuthnBrowser;

    document.addEventListener("DOMContentLoaded", ready);

    function ready() {
        document.getElementById('registerButton').addEventListener('click', begin);
        document.getElementById('loginButton').addEventListener('click', login);
    }

    function showMessage(message, isError = false) {
        const messageElement = document.getElementById('message');
        messageElement.textContent = message;
        messageElement.style.color = isError ? 'red' : 'green';
    }

    const begin = async () => {
        // Retrieve the username from the input field
        const username = document.getElementById('username').value;
        const displayname = document.getElementById('displayname').value;

        // Reset success/error messages
        elemSuccess.innerHTML = '';
        elemError.innerHTML = '';

        // GET registration options from the endpoint that calls
        // @simplewebauthn/server -> generateRegistrationOptions()
        const resp = await fetch('/api/passkey/generate-registration-options');

        let attResp;
        try {
            // Pass the options to the authenticator and wait for a response
            attResp = await startRegistration(await resp.json());
            } catch (error) {
            // Some basic error handling
            if (error.name === 'InvalidStateError') {
                elemError.innerText = 'Error: Authenticator was probably already registered by user';
            } else {
                elemError.innerText = error;
            }
            throw error;
        }

        // POST the response to the endpoint that calls
        // @simplewebauthn/server -> verifyRegistrationResponse()
        const verificationResp = await fetch('/verify-registration', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(attResp),
        });

        // Wait for the results of verification
        const verificationJSON = await verificationResp.json();

        // Show UI appropriate for the ` + "`" + `verified` + "`" + ` status
        if (verificationJSON && verificationJSON.verified) {
            elemSuccess.innerHTML = 'Success!';
        } else {
            elemError.innerHTML = ` + "`" + `Oh no, something went wrong! Response: <pre>${JSON.stringify(
                verificationJSON,
            )}</pre>` + "`" + `;
        }

//   });
//         try {
//             if (!username) {
//                 throw new Error('Please enter a username.');
//             }

//             if (!displayname) {
//                 throw new Error('Please enter a display name.');
//             }

//             // Get login options from your server. Here, we also receive the challenge.
//             const response = await fetch('/api/passkey/loginStart', {
//                 method: 'POST', headers: {'Content-Type': 'application/json'},
//                 body: JSON.stringify({username: username, displayname: displayname})
//             });
//             // Check if the login options are ok.
//             if (!response.ok) {
//                 const msg = await response.json();
//                 throw new Error('Failed to get login options from server: ' + msg);
//             }
//             // Convert the login options to JSON.
//             const options = await response.json();
//             console.log(options)

//             // This triggers the browser to display the passkey / WebAuthn modal (e.g. Face ID, Touch ID, Windows Hello).
//             // A new assertionResponse is created. This also means that the challenge has been signed.
//             const assertionResponse = await navigator.credentials.get(options.publicKey);

//             // Send assertionResponse back to server for verification.
//             const verificationResponse = await fetch('/api/passkey/loginFinish', {
//                 method: 'POST',
//                 headers: {
//                     'Content-Type': 'application/json',
//                     'Session-Key': response.headers.get('Session-Key'),
//                 },
//                 body: JSON.stringify(assertionResponse)
//             });

//             const msg = await verificationResponse.json();
//             if (verificationResponse.ok) {
//                 showMessage(msg, false);
//             } else {
//                 showMessage(msg, true);
//             }
//         } catch (error) {
//             showMessage('Error: ' + error.message, true);
//         }
    }

    const register = async () => {
        // Retrieve the username from the input field
        const username = document.getElementById('username').value;
        const displayname = document.getElementById('displayname').value;

        try {
            // Get registration options from your server. Here, we also receive the challenge.
            const response = await fetch('/api/passkey/registerStart', {
                method: 'POST', headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({username: username, displayname: displayname})
            });
            console.log(response)

            // Check if the registration options are ok.
            if (!response.ok) {
                const msg = await response.json();
                throw new Error('User already exists or failed to get registration options from server: ' + msg);
            }

            // Convert the registration options to JSON.
            let options = await response.json();
            console.log(options)

            // Set
            options = {
                publicKey: {
                    ...options,
                    user: {
                        ...options.user,
                        id: Uint8Array.from(atob(options.user.id), c => c.charCodeAt(0))
                    }
                }
            };

            // Convert options to the appropriate format for WebAuthn API
            // options.publicKey.challenge = Uint8Array.from(atob(options.publicKey.challenge), c => c.charCodeAt(0));
            // options.publicKey.user.id = Uint8Array.from(atob(options.publicKey.user.id), c => c.charCodeAt(0));
            // if (options.publicKey.excludeCredentials) {
            //     for (let cred of options.publicKey.excludeCredentials) {
            //         cred.id = Uint8Array.from(atob(cred.id), c => c.charCodeAt(0));
            //     }
            // }
            // This triggers the browser to display the passkey / WebAuthn modal (e.g. Face ID, Touch ID, Windows Hello).
            // A new attestation is created. This also means a new public-private-key pair is created.
            const attestationResponse = await navigator.credentials.create(options.publicKey);
            console.log('Attestation response:', attestationResponse);

            // Send attestationResponse back to server for verification and storage.
            const verificationResponse = await fetch('/api/passkey/registerFinish', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Session-Key': response.headers.get('Session-Key')
                },
                body: JSON.stringify(attestationResponse)
            });


            const msg = await verificationResponse.json();
            if (verificationResponse.ok) {
                showMessage(msg, false);
            } else {
                showMessage(msg, true);
            }
        } catch(error) {
            showMessage('Error: ' + error.message, true);
        }
    }
}`,
		Call:       templ.SafeScript(`__templ_pkScript_da18`),
		CallInline: templ.SafeScriptInline(`__templ_pkScript_da18`),
	}
}

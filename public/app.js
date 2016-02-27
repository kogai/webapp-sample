function onPost(target, endpoint) {
  if (!target) {
    return;
  }
  const button = target.querySelector('button');
  button.addEventListener('click', ()=> {
    const user = target.querySelector('input[type=text]').value;
    const password = target.querySelector('input[type=password]').value;
    const payload = {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({
        user,
        password,
      }),
      headers: {
        'Content-Type': 'application/json',
      },
    };

    fetch(endpoint, payload)
    .then((res)=> res.json())
    .then((json)=> {
      if (json.status !== 200) {
        
      }
      console.log(json);
    })
    .catch((error)=> console.log(error));
  });
}

const login = document.querySelector('#s-login');
const register = document.querySelector('#s-register');

onPost(login, '/login');
onPost(register, '/register');
export const loginRules = {
  email: 'required|email|min:8|max:30',
  password: ['required', 'regex:/^(?=.*[a-z])(?=.*\\d)([^\\s]){8,}$/i']
}

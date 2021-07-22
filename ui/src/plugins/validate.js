import {required, email, minLength, maxLength} from 'vuelidate/lib/validators';

class Validate
{
	install(Vue)
	{
		Vue.prototype.$v = {
			required,
			email,
			minLength,
			maxLength
		};
	}
}

export default Validate;

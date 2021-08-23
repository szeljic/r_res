import {format} from 'date-fns';

class DateTime {
	install(Vue)
	{
		Vue.prototype.$format = format;
	}
}

export default DateTime;

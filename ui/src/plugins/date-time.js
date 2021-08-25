import {format, parseISO} from 'date-fns';

class DateTime {
	install(Vue)
	{
		const dateFormat = 'dd. MM. yyyy.',
			timeFormat = 'hh:mm:ss';

		Vue.prototype.$dateFormat = format;
		Vue.prototype.$dateParseISO = parseISO;

		Vue.prototype.$dateFormatL18n = function (v, hms = false)
		{
			return format(v, dateFormat + (hms ? ' ' + timeFormat : ''));
		};
	}
}

export default DateTime;

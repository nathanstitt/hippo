import "./home.scss";
import $ from 'domtastic';

$(document).ready(() => {
    $('input, textarea').on('keyup blur focus', function (e) {


        var $this = $(e.target),
            label = $(e.target.parentElement.querySelector('label'));
        if (e.type === 'keyup') {
            if ($this.val() === '') {
                label.removeClass('active highlight');
            } else {
                label.addClass('active highlight');
            }
        } else if (e.type === 'blur') {
            if( $this.val() === '' ) {
                label.removeClass('active highlight');
            } else {
                label.removeClass('highlight');
            }
        } else if (e.type === 'focus') {

            if( $this.val() === '' ) {
                label.removeClass('highlight');
            }
            else if( $this.val() !== '' ) {
                label.addClass('highlight');
            }
        }
    });


    $('.tab a').on('click', function (e) {
        e.preventDefault();
        $(this).parent().addClass('active');
        $(this).parent().siblings().removeClass('active');
        var target = $(this).attr('href');
        $('.tab-content > div').forEach((d) => {
            if ($.matches(d, target)) {
                $(d).removeClass('hidden')
            } else {
                $(d).addClass('hidden')
            }
        })
    });
});

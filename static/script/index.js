'use strict';

document.addEventListener('DOMContentLoaded', () => {

    const imageForm = {};  //imageForm храненит информации о загруженных изображениях в форму.

    const form = document.querySelector('#adminPost'),              //получает ссылку
     inputsText = form.querySelectorAll('.admin-form__input-text'), //получает массив
     inputsTextErrorClass = 'admin-form__input-text_error',         //содержит ошибку валидации
     hiddenClass = 'hidden';               //содержит класс для скрытия элементов формы от пользователя
     //элементы DOM формы

    const inputTextAndPreviewElemId = {
        postTitle: ['previewTitleArticle', 'previewTitlePostCard'],
        postShortDescr: ['previewDescrArticle', 'previewDescrPostCard'],
        postAuthorName: ['previewAuthorNamePostCard'],
        postPublishDate: ['previewDatePostCard']
    }; //содержит информацию о соответствии полей ввода текста и соответствующих им элементов которые будут обновляться

    const componentsImageObj = {
        inputsImage: {
            postAdminPhoto: {
                inputWrapper: 'adminFormAuthorPhoto',
                previewsObj: ['previewAuthorPhotoInput', 'previewAuthorPhotoPostCard'],
                deleteUploadBtnBoolean: true,
                deleteImageInfo: false
            },
            postBigImage: {
                inputWrapper: 'adminFormBigImage',
                previewsObj: ['previewBigImageInput', 'previewBigImageArticle'],
                deleteUploadBtnBoolean: false,
                deleteImageInfo: true
            },
            postSmallImage: {
                inputWrapper: 'adminFormSmallImage',
                previewsObj: ['previewSmallImageInput', 'previewSmallImagePostCard'],
                deleteUploadBtnBoolean: false,
                deleteImageInfo: true
            }
        },

        buttons: {
            btnImageUploadSelector: '.admin-form__input-upload',
            removeAndUploadNewBtnsWrapper: '.admin-form__input-btn-wrapper',
            imageInfo: '.admin-form__image-info',
            removeBtnSelector: '.admin-form__input-remove'
        }

    }; //содержит информацию об элементах формы, связанных с загрузкой изображений,
    //таких как кнопки, информация об изображении и т.д.

    const uploadImageInFormData = async (imageForm, input) => {
        let reader = new FileReader();

        if (input.files[0]) {
            reader.readAsDataURL(input.files[0]);
            reader.onloadend  =  () => {
                imageForm[input.getAttribute('name')] = {
                    imageInBase64: String(reader.result),
                    nameFile: input.files[0].name
                };
            };
        } else {
            imageForm[input.getAttribute('name')] = {};
        }
    }; //загружает изображения в imageForm

    const workWithImageInput = (componentsImageObj) => {

        const inputsImageObj =  componentsImageObj.inputsImage;

        const updatePreviews = (arrPreviews, style) =>{
            arrPreviews.forEach(preview => {
                 document.querySelector(`#${preview}`).style.background = `${style}`;
            });
        };

         const showAndHideUploadNewAndRemoveBtns = (show, imageObj, wrapper, wrapperBtns) =>  {
            if (show) {
                if (imageObj.deleteUploadBtnBoolean) {
                    wrapper.querySelector(componentsImageObj.buttons.btnImageUploadSelector).classList.add(hiddenClass);
                }

                if (imageObj.deleteImageInfo) {
                    wrapper.querySelector(componentsImageObj.buttons.imageInfo).classList.add(hiddenClass);
                }
                wrapperBtns.classList.remove(hiddenClass);
            } else {
                if (imageObj.deleteUploadBtnBoolean) {
                    wrapper.querySelector(componentsImageObj.buttons.btnImageUploadSelector).classList.remove(hiddenClass);
                }

                if (imageObj.deleteImageInfo) {
                    wrapper.querySelector(componentsImageObj.buttons.imageInfo).classList.remove(hiddenClass);
                }
                wrapperBtns.classList.add(hiddenClass);
            }
        };

        for (let key in inputsImageObj) {
            let input = document.querySelector(`#${key}`),
            wrapper = document.querySelector(`#${inputsImageObj[key].inputWrapper}`),
            wrapperBtns = wrapper.querySelector(componentsImageObj.buttons.removeAndUploadNewBtnsWrapper),
            removeBtn = wrapper.querySelector(componentsImageObj.buttons.removeBtnSelector);

            input.addEventListener('input', () => {

                if (input.files[0]) {
                    updatePreviews(inputsImageObj[key].previewsObj,
                        `url('${window.URL.createObjectURL(input.files[0])}') center center/cover no-repeat`);

                    showAndHideUploadNewAndRemoveBtns(true, inputsImageObj[key], wrapper, wrapperBtns);

                    uploadImageInFormData(imageForm, input);
                }

            });

            removeBtn.addEventListener('click', (event) => {
                event.preventDefault();

                const attrFor = event.target.getAttribute('for'),
                    input = document.querySelector(`#${attrFor}`);

                input.value = '';

                updatePreviews(inputsImageObj[key].previewsObj, '');

                showAndHideUploadNewAndRemoveBtns(false, inputsImageObj[key], wrapper, wrapperBtns);

                uploadImageInFormData(imageForm, input);

            });

        }
    }; //обрабатывает события, связанные с загрузкой и удалением изображений, а также обновляет соответствующие
       // элементы предпросмотра и изменяет состояние кнопок.

    workWithImageInput(componentsImageObj);

    const activateErrorInputText = (input, inputsTextErrorClass, hiddenClass) => {
        input.classList.add(inputsTextErrorClass);
        input.nextElementSibling.classList.remove(hiddenClass);
    }; //используется для стилизации поля ввода и показа сообщения об ошибке в случае
       //некорректного заполнения поля ввода.

    const deactivateErrorInputText = (input, inputsTextErrorClass, hiddenClass) => {
        input.classList.remove(inputsTextErrorClass);
        input.nextElementSibling.classList.add(hiddenClass);
    }; //используется для удаления стилизации поля ввода и скрытия сообщения об ошибке
       //после исправления ошибки.

    const validateForm = (inputsText) => {
        return validateInputsText(inputsText);
    };  //проверка корректности введенных значений в поля формы.

    const validateInputsText = (inputs) => {
        let validate = true;
        inputs.forEach(input => {
            if(input.value === '' || input.value === null ||  input.value === undefined) {
                validate = false;
                activateErrorInputText(input, inputsTextErrorClass, hiddenClass); //показывает сообщение об ошибке
            }
        });

        return validate;
    }; //validateInputsText используются для проверки корректности заполнения полей ввода текста.

    const showPreviewTextInputs = (inputTextAndPreviewElemId) => {
        let elem, previewElem;
        for(let elemId in inputTextAndPreviewElemId) {
            elem = document.querySelector(`#${elemId}`);
            inputTextAndPreviewElemId[elemId].forEach(idPreview => {
                previewElem = document.querySelector(`#${idPreview}`);
                if (elem.value != '') {
                    previewElem.innerHTML = elem.value;
                } else {
                    previewElem.innerHTML = previewElem.getAttribute('data-default-value');
                }
            });
        }
    }; //обновляет элементы предпросмотра с учетом введенных данных в поля ввода.

    inputsText.forEach(input => {
        input.addEventListener('input', () => {
            showPreviewTextInputs(inputTextAndPreviewElemId);
            if(input.value === '' || input.value === null ||  input.value === undefined) {
                activateErrorInputText(input, inputsTextErrorClass, hiddenClass);
            } else {
                deactivateErrorInputText (input, inputsTextErrorClass, hiddenClass);
            }
        });
    });


    form.addEventListener('submit', (event) => {
        event.preventDefault();
        const formData = new FormData(form),
              newFormData = Object.fromEntries(formData),
              infoComponent = form.querySelector('.admin-form__info');

        const infoAboutStateSubmitForm = {
            success: {
                infoClass: 'admin-form__success',
                text: 'Publish Complete!'
            },
            error: {
                infoClass: 'admin-form__error',
                text: 'Whoops! Some fields need your attention :o',
            }
        };  //submit на форме формирует объект formData, содержащий данные формы, добавляет данные об
            //изображениях из imageForm и выводит информацию о результате отправки формы в консоль.




        if (validateForm(inputsText)) {
            for (let key in imageForm) {
                newFormData[key] = imageForm[key];
            }

            infoComponent.classList.remove(infoAboutStateSubmitForm.error.infoClass);
            infoComponent.classList.add(infoAboutStateSubmitForm.success.infoClass);
            infoComponent.classList.remove(hiddenClass);
            infoComponent.innerHTML = infoAboutStateSubmitForm.success.text;

            console.log(newFormData);
            console.log(JSON.stringify(newFormData));
        } else {

            infoComponent.classList.remove(infoAboutStateSubmitForm.success.infoClass);
            infoComponent.classList.add(infoAboutStateSubmitForm.error.infoClass);
            infoComponent.classList.remove(hiddenClass);
            infoComponent.innerHTML = infoAboutStateSubmitForm.error.text;

            console.error('Ошибка!');
        }

    });


});
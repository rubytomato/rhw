# �T�v

Go��Web Application Framework�ł���Revel���g���ĊJ�������ȒP�Ȍ����A�v���P�[�V�����ł��B


**��**

* Go 1.4.2
* Revel 0.12.0


## ���O����

Revel�̃C���X�g�[��

```
> go get -u -v github.com/revel/revel
> go get -u -v github.com/revel/cmd/revel
```


## �A�v���P�[�V�����̎��s

�A�v���P�[�V�����̃��[�g�f�B���N�g���ŉ��L��revel run�R�}���h�����s���܂��B

```
> revel.exe run rhw
```

�A�v���P�[�V�������N�������牺�L��URL�ɃA�N�Z�X����ƃy�[�W�����邱�Ƃ��ł��܂��B
�A�v���P�[�V�����̒�~�̓R���\�[������Ctrl + C�ōs���܂��B

http://localhost:9000


## �A�v���P�[�V�����̃r���h

revel build�R�}���h�̓A�v���P�[�V�����̎��s�\�ȃo�C�i���t�@�C���𐶐����܂��B(Windows����exe�t�@�C��)
�r���h������O�Ƀr���h�Ŏg�p����e���|�����[�t�H���_���K�v�ł��B���̗�ł�C:/tmp�ɍ쐬���܂����B

```
> revel.exe build rhw "c:/tmp/rhw"
```

�r���h����������ƃe���|�����[�t�H���_�ɉ��L�̃t�@�C������������܂��B

```
> dir /b
run.bat
run.sh
src
rhw.exe
```

run.bat�t�@�C�������s�����web�A�v���P�[�V�������N�����܂��B

```
> run.bat
```

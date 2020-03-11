
double-float.out:     file format elf64-littleriscv


Disassembly of section .text:

00000000000100b0 <register_fini>:
   100b0:	ffff0797          	auipc	a5,0xffff0
   100b4:	f5078793          	addi	a5,a5,-176 # 0 <register_fini-0x100b0>
   100b8:	cb89                	beqz	a5,100ca <register_fini+0x1a>
   100ba:	00000517          	auipc	a0,0x0
   100be:	2f650513          	addi	a0,a0,758 # 103b0 <__libc_fini_array>
   100c2:	00000317          	auipc	t1,0x0
   100c6:	2b230067          	jr	690(t1) # 10374 <atexit>
   100ca:	8082                	ret

00000000000100cc <_start>:
   100cc:	00002197          	auipc	gp,0x2
   100d0:	ddc18193          	addi	gp,gp,-548 # 11ea8 <__global_pointer$>
   100d4:	00002517          	auipc	a0,0x2
   100d8:	dac50513          	addi	a0,a0,-596 # 11e80 <_edata>
   100dc:	00002617          	auipc	a2,0x2
   100e0:	ddc60613          	addi	a2,a2,-548 # 11eb8 <__BSS_END__>
   100e4:	8e09                	sub	a2,a2,a0
   100e6:	4581                	li	a1,0
   100e8:	00000097          	auipc	ra,0x0
   100ec:	36a080e7          	jalr	874(ra) # 10452 <memset>
   100f0:	00000517          	auipc	a0,0x0
   100f4:	2c050513          	addi	a0,a0,704 # 103b0 <__libc_fini_array>
   100f8:	00000097          	auipc	ra,0x0
   100fc:	27c080e7          	jalr	636(ra) # 10374 <atexit>
   10100:	00000097          	auipc	ra,0x0
   10104:	2e8080e7          	jalr	744(ra) # 103e8 <__libc_init_array>
   10108:	4502                	lw	a0,0(sp)
   1010a:	002c                	addi	a1,sp,8
   1010c:	4601                	li	a2,0
   1010e:	00000097          	auipc	ra,0x0
   10112:	07a080e7          	jalr	122(ra) # 10188 <main>
   10116:	00000317          	auipc	t1,0x0
   1011a:	26e30067          	jr	622(t1) # 10384 <exit>

000000000001011e <__do_global_dtors_aux>:
   1011e:	00002797          	auipc	a5,0x2
   10122:	d627c783          	lbu	a5,-670(a5) # 11e80 <_edata>
   10126:	ef95                	bnez	a5,10162 <__do_global_dtors_aux+0x44>
   10128:	ffff0797          	auipc	a5,0xffff0
   1012c:	ed878793          	addi	a5,a5,-296 # 0 <register_fini-0x100b0>
   10130:	c39d                	beqz	a5,10156 <__do_global_dtors_aux+0x38>
   10132:	1141                	addi	sp,sp,-16
   10134:	00001517          	auipc	a0,0x1
   10138:	55450513          	addi	a0,a0,1364 # 11688 <__FRAME_END__>
   1013c:	e406                	sd	ra,8(sp)
   1013e:	00000097          	auipc	ra,0x0
   10142:	000000e7          	jalr	zero # 0 <register_fini-0x100b0>
   10146:	60a2                	ld	ra,8(sp)
   10148:	4785                	li	a5,1
   1014a:	00002717          	auipc	a4,0x2
   1014e:	d2f70b23          	sb	a5,-714(a4) # 11e80 <_edata>
   10152:	0141                	addi	sp,sp,16
   10154:	8082                	ret
   10156:	4785                	li	a5,1
   10158:	00002717          	auipc	a4,0x2
   1015c:	d2f70423          	sb	a5,-728(a4) # 11e80 <_edata>
   10160:	8082                	ret
   10162:	8082                	ret

0000000000010164 <frame_dummy>:
   10164:	ffff0797          	auipc	a5,0xffff0
   10168:	e9c78793          	addi	a5,a5,-356 # 0 <register_fini-0x100b0>
   1016c:	cf89                	beqz	a5,10186 <frame_dummy+0x22>
   1016e:	00002597          	auipc	a1,0x2
   10172:	d1a58593          	addi	a1,a1,-742 # 11e88 <object.5475>
   10176:	00001517          	auipc	a0,0x1
   1017a:	51250513          	addi	a0,a0,1298 # 11688 <__FRAME_END__>
   1017e:	00000317          	auipc	t1,0x0
   10182:	00000067          	jr	zero # 0 <register_fini-0x100b0>
   10186:	8082                	ret

0000000000010188 <main>:
 
float result_float[10]={1.1,2.2,3.3,4.4,5.5,6.6,7.7,8.8,9.9,11.0};
//result float : 5.5,5.5,5.5,5.5,5.5,1.1,1.1,1.1,1.1,1.1

int main()
{
   10188:	fe010113          	addi	sp,sp,-32
   1018c:	00813c23          	sd	s0,24(sp)
   10190:	02010413          	addi	s0,sp,32
	int i=0;
   10194:	fe042623          	sw	zero,-20(s0)
	for(i=4;i>=0;i--)
   10198:	00400793          	li	a5,4
   1019c:	fef42623          	sw	a5,-20(s0)
   101a0:	0580006f          	j	101f8 <main+0x70>
		result_double[i]=result_double[i]+i*1.1;
   101a4:	000117b7          	lui	a5,0x11
   101a8:	6a878713          	addi	a4,a5,1704 # 116a8 <__fini_array_end>
   101ac:	fec42783          	lw	a5,-20(s0)
   101b0:	00379793          	slli	a5,a5,0x3
   101b4:	00f707b3          	add	a5,a4,a5
   101b8:	0007b707          	fld	fa4,0(a5)
   101bc:	fec42783          	lw	a5,-20(s0)
   101c0:	d20786d3          	fcvt.d.w	fa3,a5
   101c4:	000107b7          	lui	a5,0x10
   101c8:	6807b787          	fld	fa5,1664(a5) # 10680 <__errno+0x12>
   101cc:	12f6f7d3          	fmul.d	fa5,fa3,fa5
   101d0:	02f777d3          	fadd.d	fa5,fa4,fa5
   101d4:	000117b7          	lui	a5,0x11
   101d8:	6a878713          	addi	a4,a5,1704 # 116a8 <__fini_array_end>
   101dc:	fec42783          	lw	a5,-20(s0)
   101e0:	00379793          	slli	a5,a5,0x3
   101e4:	00f707b3          	add	a5,a4,a5
   101e8:	00f7b027          	fsd	fa5,0(a5)
	for(i=4;i>=0;i--)
   101ec:	fec42783          	lw	a5,-20(s0)
   101f0:	fff7879b          	addiw	a5,a5,-1
   101f4:	fef42623          	sw	a5,-20(s0)
   101f8:	fec42783          	lw	a5,-20(s0)
   101fc:	0007879b          	sext.w	a5,a5
   10200:	fa07d2e3          	bgez	a5,101a4 <main+0x1c>
	for(i=5;i<10;i++)
   10204:	00500793          	li	a5,5
   10208:	fef42623          	sw	a5,-20(s0)
   1020c:	0580006f          	j	10264 <main+0xdc>
		result_double[i]=result_double[i]-i*1.1;
   10210:	000117b7          	lui	a5,0x11
   10214:	6a878713          	addi	a4,a5,1704 # 116a8 <__fini_array_end>
   10218:	fec42783          	lw	a5,-20(s0)
   1021c:	00379793          	slli	a5,a5,0x3
   10220:	00f707b3          	add	a5,a4,a5
   10224:	0007b707          	fld	fa4,0(a5)
   10228:	fec42783          	lw	a5,-20(s0)
   1022c:	d20786d3          	fcvt.d.w	fa3,a5
   10230:	000107b7          	lui	a5,0x10
   10234:	6807b787          	fld	fa5,1664(a5) # 10680 <__errno+0x12>
   10238:	12f6f7d3          	fmul.d	fa5,fa3,fa5
   1023c:	0af777d3          	fsub.d	fa5,fa4,fa5
   10240:	000117b7          	lui	a5,0x11
   10244:	6a878713          	addi	a4,a5,1704 # 116a8 <__fini_array_end>
   10248:	fec42783          	lw	a5,-20(s0)
   1024c:	00379793          	slli	a5,a5,0x3
   10250:	00f707b3          	add	a5,a4,a5
   10254:	00f7b027          	fsd	fa5,0(a5)
	for(i=5;i<10;i++)
   10258:	fec42783          	lw	a5,-20(s0)
   1025c:	0017879b          	addiw	a5,a5,1
   10260:	fef42623          	sw	a5,-20(s0)
   10264:	fec42783          	lw	a5,-20(s0)
   10268:	0007871b          	sext.w	a4,a5
   1026c:	00900793          	li	a5,9
   10270:	fae7d0e3          	bge	a5,a4,10210 <main+0x88>
	
	for(i=4;i>=0;i--)
   10274:	00400793          	li	a5,4
   10278:	fef42623          	sw	a5,-20(s0)
   1027c:	0600006f          	j	102dc <main+0x154>
		result_float[i]=result_float[i]+i*1.1;
   10280:	000117b7          	lui	a5,0x11
   10284:	6f878713          	addi	a4,a5,1784 # 116f8 <result_float>
   10288:	fec42783          	lw	a5,-20(s0)
   1028c:	00279793          	slli	a5,a5,0x2
   10290:	00f707b3          	add	a5,a4,a5
   10294:	0007a787          	flw	fa5,0(a5)
   10298:	42078753          	fcvt.d.s	fa4,fa5
   1029c:	fec42783          	lw	a5,-20(s0)
   102a0:	d20786d3          	fcvt.d.w	fa3,a5
   102a4:	000107b7          	lui	a5,0x10
   102a8:	6807b787          	fld	fa5,1664(a5) # 10680 <__errno+0x12>
   102ac:	12f6f7d3          	fmul.d	fa5,fa3,fa5
   102b0:	02f777d3          	fadd.d	fa5,fa4,fa5
   102b4:	4017f7d3          	fcvt.s.d	fa5,fa5
   102b8:	000117b7          	lui	a5,0x11
   102bc:	6f878713          	addi	a4,a5,1784 # 116f8 <result_float>
   102c0:	fec42783          	lw	a5,-20(s0)
   102c4:	00279793          	slli	a5,a5,0x2
   102c8:	00f707b3          	add	a5,a4,a5
   102cc:	00f7a027          	fsw	fa5,0(a5)
	for(i=4;i>=0;i--)
   102d0:	fec42783          	lw	a5,-20(s0)
   102d4:	fff7879b          	addiw	a5,a5,-1
   102d8:	fef42623          	sw	a5,-20(s0)
   102dc:	fec42783          	lw	a5,-20(s0)
   102e0:	0007879b          	sext.w	a5,a5
   102e4:	f807dee3          	bgez	a5,10280 <main+0xf8>
	for(i=5;i<10;i++)
   102e8:	00500793          	li	a5,5
   102ec:	fef42623          	sw	a5,-20(s0)
   102f0:	0600006f          	j	10350 <main+0x1c8>
		result_float[i]=result_float[i]-i*1.1;
   102f4:	000117b7          	lui	a5,0x11
   102f8:	6f878713          	addi	a4,a5,1784 # 116f8 <result_float>
   102fc:	fec42783          	lw	a5,-20(s0)
   10300:	00279793          	slli	a5,a5,0x2
   10304:	00f707b3          	add	a5,a4,a5
   10308:	0007a787          	flw	fa5,0(a5)
   1030c:	42078753          	fcvt.d.s	fa4,fa5
   10310:	fec42783          	lw	a5,-20(s0)
   10314:	d20786d3          	fcvt.d.w	fa3,a5
   10318:	000107b7          	lui	a5,0x10
   1031c:	6807b787          	fld	fa5,1664(a5) # 10680 <__errno+0x12>
   10320:	12f6f7d3          	fmul.d	fa5,fa3,fa5
   10324:	0af777d3          	fsub.d	fa5,fa4,fa5
   10328:	4017f7d3          	fcvt.s.d	fa5,fa5
   1032c:	000117b7          	lui	a5,0x11
   10330:	6f878713          	addi	a4,a5,1784 # 116f8 <result_float>
   10334:	fec42783          	lw	a5,-20(s0)
   10338:	00279793          	slli	a5,a5,0x2
   1033c:	00f707b3          	add	a5,a4,a5
   10340:	00f7a027          	fsw	fa5,0(a5)
	for(i=5;i<10;i++)
   10344:	fec42783          	lw	a5,-20(s0)
   10348:	0017879b          	addiw	a5,a5,1
   1034c:	fef42623          	sw	a5,-20(s0)
   10350:	fec42783          	lw	a5,-20(s0)
   10354:	0007871b          	sext.w	a4,a5
   10358:	00900793          	li	a5,9
   1035c:	f8e7dce3          	bge	a5,a4,102f4 <main+0x16c>
	
	return 0;
   10360:	00000793          	li	a5,0
   10364:	00078513          	mv	a0,a5
   10368:	01813403          	ld	s0,24(sp)
   1036c:	02010113          	addi	sp,sp,32
   10370:	00008067          	ret

0000000000010374 <atexit>:
   10374:	85aa                	mv	a1,a0
   10376:	4681                	li	a3,0
   10378:	4601                	li	a2,0
   1037a:	4501                	li	a0,0
   1037c:	00000317          	auipc	t1,0x0
   10380:	18030067          	jr	384(t1) # 104fc <__register_exitproc>

0000000000010384 <exit>:
   10384:	1141                	addi	sp,sp,-16
   10386:	4581                	li	a1,0
   10388:	e022                	sd	s0,0(sp)
   1038a:	e406                	sd	ra,8(sp)
   1038c:	842a                	mv	s0,a0
   1038e:	00000097          	auipc	ra,0x0
   10392:	1ea080e7          	jalr	490(ra) # 10578 <__call_exitprocs>
   10396:	00002797          	auipc	a5,0x2
   1039a:	ad278793          	addi	a5,a5,-1326 # 11e68 <_global_impure_ptr>
   1039e:	6388                	ld	a0,0(a5)
   103a0:	6d3c                	ld	a5,88(a0)
   103a2:	c391                	beqz	a5,103a6 <exit+0x22>
   103a4:	9782                	jalr	a5
   103a6:	8522                	mv	a0,s0
   103a8:	00000097          	auipc	ra,0x0
   103ac:	296080e7          	jalr	662(ra) # 1063e <_exit>

00000000000103b0 <__libc_fini_array>:
   103b0:	1101                	addi	sp,sp,-32
   103b2:	e822                	sd	s0,16(sp)
   103b4:	00001797          	auipc	a5,0x1
   103b8:	2f478793          	addi	a5,a5,756 # 116a8 <__fini_array_end>
   103bc:	00001417          	auipc	s0,0x1
   103c0:	2e440413          	addi	s0,s0,740 # 116a0 <__init_array_end>
   103c4:	8f81                	sub	a5,a5,s0
   103c6:	e426                	sd	s1,8(sp)
   103c8:	ec06                	sd	ra,24(sp)
   103ca:	4037d493          	srai	s1,a5,0x3
   103ce:	c881                	beqz	s1,103de <__libc_fini_array+0x2e>
   103d0:	17e1                	addi	a5,a5,-8
   103d2:	943e                	add	s0,s0,a5
   103d4:	601c                	ld	a5,0(s0)
   103d6:	14fd                	addi	s1,s1,-1
   103d8:	1461                	addi	s0,s0,-8
   103da:	9782                	jalr	a5
   103dc:	fce5                	bnez	s1,103d4 <__libc_fini_array+0x24>
   103de:	60e2                	ld	ra,24(sp)
   103e0:	6442                	ld	s0,16(sp)
   103e2:	64a2                	ld	s1,8(sp)
   103e4:	6105                	addi	sp,sp,32
   103e6:	8082                	ret

00000000000103e8 <__libc_init_array>:
   103e8:	1101                	addi	sp,sp,-32
   103ea:	e822                	sd	s0,16(sp)
   103ec:	e04a                	sd	s2,0(sp)
   103ee:	00001417          	auipc	s0,0x1
   103f2:	29e40413          	addi	s0,s0,670 # 1168c <__preinit_array_end>
   103f6:	00001917          	auipc	s2,0x1
   103fa:	29690913          	addi	s2,s2,662 # 1168c <__preinit_array_end>
   103fe:	40890933          	sub	s2,s2,s0
   10402:	ec06                	sd	ra,24(sp)
   10404:	e426                	sd	s1,8(sp)
   10406:	40395913          	srai	s2,s2,0x3
   1040a:	00090963          	beqz	s2,1041c <__libc_init_array+0x34>
   1040e:	4481                	li	s1,0
   10410:	601c                	ld	a5,0(s0)
   10412:	0485                	addi	s1,s1,1
   10414:	0421                	addi	s0,s0,8
   10416:	9782                	jalr	a5
   10418:	fe991ce3          	bne	s2,s1,10410 <__libc_init_array+0x28>
   1041c:	00001417          	auipc	s0,0x1
   10420:	27440413          	addi	s0,s0,628 # 11690 <__init_array_start>
   10424:	00001917          	auipc	s2,0x1
   10428:	27c90913          	addi	s2,s2,636 # 116a0 <__init_array_end>
   1042c:	40890933          	sub	s2,s2,s0
   10430:	40395913          	srai	s2,s2,0x3
   10434:	00090963          	beqz	s2,10446 <__libc_init_array+0x5e>
   10438:	4481                	li	s1,0
   1043a:	601c                	ld	a5,0(s0)
   1043c:	0485                	addi	s1,s1,1
   1043e:	0421                	addi	s0,s0,8
   10440:	9782                	jalr	a5
   10442:	fe991ce3          	bne	s2,s1,1043a <__libc_init_array+0x52>
   10446:	60e2                	ld	ra,24(sp)
   10448:	6442                	ld	s0,16(sp)
   1044a:	64a2                	ld	s1,8(sp)
   1044c:	6902                	ld	s2,0(sp)
   1044e:	6105                	addi	sp,sp,32
   10450:	8082                	ret

0000000000010452 <memset>:
   10452:	433d                	li	t1,15
   10454:	872a                	mv	a4,a0
   10456:	02c37163          	bgeu	t1,a2,10478 <memset+0x26>
   1045a:	00f77793          	andi	a5,a4,15
   1045e:	e3c1                	bnez	a5,104de <memset+0x8c>
   10460:	e1bd                	bnez	a1,104c6 <memset+0x74>
   10462:	ff067693          	andi	a3,a2,-16
   10466:	8a3d                	andi	a2,a2,15
   10468:	96ba                	add	a3,a3,a4
   1046a:	e30c                	sd	a1,0(a4)
   1046c:	e70c                	sd	a1,8(a4)
   1046e:	0741                	addi	a4,a4,16
   10470:	fed76de3          	bltu	a4,a3,1046a <memset+0x18>
   10474:	e211                	bnez	a2,10478 <memset+0x26>
   10476:	8082                	ret
   10478:	40c306b3          	sub	a3,t1,a2
   1047c:	068a                	slli	a3,a3,0x2
   1047e:	00000297          	auipc	t0,0x0
   10482:	9696                	add	a3,a3,t0
   10484:	00a68067          	jr	10(a3)
   10488:	00b70723          	sb	a1,14(a4)
   1048c:	00b706a3          	sb	a1,13(a4)
   10490:	00b70623          	sb	a1,12(a4)
   10494:	00b705a3          	sb	a1,11(a4)
   10498:	00b70523          	sb	a1,10(a4)
   1049c:	00b704a3          	sb	a1,9(a4)
   104a0:	00b70423          	sb	a1,8(a4)
   104a4:	00b703a3          	sb	a1,7(a4)
   104a8:	00b70323          	sb	a1,6(a4)
   104ac:	00b702a3          	sb	a1,5(a4)
   104b0:	00b70223          	sb	a1,4(a4)
   104b4:	00b701a3          	sb	a1,3(a4)
   104b8:	00b70123          	sb	a1,2(a4)
   104bc:	00b700a3          	sb	a1,1(a4)
   104c0:	00b70023          	sb	a1,0(a4)
   104c4:	8082                	ret
   104c6:	0ff5f593          	andi	a1,a1,255
   104ca:	00859693          	slli	a3,a1,0x8
   104ce:	8dd5                	or	a1,a1,a3
   104d0:	01059693          	slli	a3,a1,0x10
   104d4:	8dd5                	or	a1,a1,a3
   104d6:	02059693          	slli	a3,a1,0x20
   104da:	8dd5                	or	a1,a1,a3
   104dc:	b759                	j	10462 <memset+0x10>
   104de:	00279693          	slli	a3,a5,0x2
   104e2:	00000297          	auipc	t0,0x0
   104e6:	9696                	add	a3,a3,t0
   104e8:	8286                	mv	t0,ra
   104ea:	fa2680e7          	jalr	-94(a3)
   104ee:	8096                	mv	ra,t0
   104f0:	17c1                	addi	a5,a5,-16
   104f2:	8f1d                	sub	a4,a4,a5
   104f4:	963e                	add	a2,a2,a5
   104f6:	f8c371e3          	bgeu	t1,a2,10478 <memset+0x26>
   104fa:	b79d                	j	10460 <memset+0xe>

00000000000104fc <__register_exitproc>:
   104fc:	00002797          	auipc	a5,0x2
   10500:	96c78793          	addi	a5,a5,-1684 # 11e68 <_global_impure_ptr>
   10504:	6398                	ld	a4,0(a5)
   10506:	1f873783          	ld	a5,504(a4)
   1050a:	c3b1                	beqz	a5,1054e <__register_exitproc+0x52>
   1050c:	4798                	lw	a4,8(a5)
   1050e:	487d                	li	a6,31
   10510:	06e84263          	blt	a6,a4,10574 <__register_exitproc+0x78>
   10514:	c505                	beqz	a0,1053c <__register_exitproc+0x40>
   10516:	00371813          	slli	a6,a4,0x3
   1051a:	983e                	add	a6,a6,a5
   1051c:	10c83823          	sd	a2,272(a6)
   10520:	3107a883          	lw	a7,784(a5)
   10524:	4605                	li	a2,1
   10526:	00e6163b          	sllw	a2,a2,a4
   1052a:	00c8e8b3          	or	a7,a7,a2
   1052e:	3117a823          	sw	a7,784(a5)
   10532:	20d83823          	sd	a3,528(a6)
   10536:	4689                	li	a3,2
   10538:	02d50063          	beq	a0,a3,10558 <__register_exitproc+0x5c>
   1053c:	00270693          	addi	a3,a4,2
   10540:	068e                	slli	a3,a3,0x3
   10542:	2705                	addiw	a4,a4,1
   10544:	c798                	sw	a4,8(a5)
   10546:	97b6                	add	a5,a5,a3
   10548:	e38c                	sd	a1,0(a5)
   1054a:	4501                	li	a0,0
   1054c:	8082                	ret
   1054e:	20070793          	addi	a5,a4,512
   10552:	1ef73c23          	sd	a5,504(a4)
   10556:	bf5d                	j	1050c <__register_exitproc+0x10>
   10558:	3147a683          	lw	a3,788(a5)
   1055c:	4501                	li	a0,0
   1055e:	8e55                	or	a2,a2,a3
   10560:	00270693          	addi	a3,a4,2
   10564:	068e                	slli	a3,a3,0x3
   10566:	2705                	addiw	a4,a4,1
   10568:	30c7aa23          	sw	a2,788(a5)
   1056c:	c798                	sw	a4,8(a5)
   1056e:	97b6                	add	a5,a5,a3
   10570:	e38c                	sd	a1,0(a5)
   10572:	8082                	ret
   10574:	557d                	li	a0,-1
   10576:	8082                	ret

0000000000010578 <__call_exitprocs>:
   10578:	715d                	addi	sp,sp,-80
   1057a:	00002797          	auipc	a5,0x2
   1057e:	8ee78793          	addi	a5,a5,-1810 # 11e68 <_global_impure_ptr>
   10582:	e062                	sd	s8,0(sp)
   10584:	0007bc03          	ld	s8,0(a5)
   10588:	f44e                	sd	s3,40(sp)
   1058a:	f052                	sd	s4,32(sp)
   1058c:	ec56                	sd	s5,24(sp)
   1058e:	e85a                	sd	s6,16(sp)
   10590:	e486                	sd	ra,72(sp)
   10592:	e0a2                	sd	s0,64(sp)
   10594:	fc26                	sd	s1,56(sp)
   10596:	f84a                	sd	s2,48(sp)
   10598:	e45e                	sd	s7,8(sp)
   1059a:	8aaa                	mv	s5,a0
   1059c:	8b2e                	mv	s6,a1
   1059e:	4a05                	li	s4,1
   105a0:	59fd                	li	s3,-1
   105a2:	1f8c3903          	ld	s2,504(s8)
   105a6:	02090463          	beqz	s2,105ce <__call_exitprocs+0x56>
   105aa:	00892483          	lw	s1,8(s2)
   105ae:	fff4841b          	addiw	s0,s1,-1
   105b2:	00044e63          	bltz	s0,105ce <__call_exitprocs+0x56>
   105b6:	048e                	slli	s1,s1,0x3
   105b8:	94ca                	add	s1,s1,s2
   105ba:	020b0663          	beqz	s6,105e6 <__call_exitprocs+0x6e>
   105be:	2084b783          	ld	a5,520(s1)
   105c2:	03678263          	beq	a5,s6,105e6 <__call_exitprocs+0x6e>
   105c6:	347d                	addiw	s0,s0,-1
   105c8:	14e1                	addi	s1,s1,-8
   105ca:	ff3418e3          	bne	s0,s3,105ba <__call_exitprocs+0x42>
   105ce:	60a6                	ld	ra,72(sp)
   105d0:	6406                	ld	s0,64(sp)
   105d2:	74e2                	ld	s1,56(sp)
   105d4:	7942                	ld	s2,48(sp)
   105d6:	79a2                	ld	s3,40(sp)
   105d8:	7a02                	ld	s4,32(sp)
   105da:	6ae2                	ld	s5,24(sp)
   105dc:	6b42                	ld	s6,16(sp)
   105de:	6ba2                	ld	s7,8(sp)
   105e0:	6c02                	ld	s8,0(sp)
   105e2:	6161                	addi	sp,sp,80
   105e4:	8082                	ret
   105e6:	00892783          	lw	a5,8(s2)
   105ea:	6498                	ld	a4,8(s1)
   105ec:	37fd                	addiw	a5,a5,-1
   105ee:	04878263          	beq	a5,s0,10632 <__call_exitprocs+0xba>
   105f2:	0004b423          	sd	zero,8(s1)
   105f6:	db61                	beqz	a4,105c6 <__call_exitprocs+0x4e>
   105f8:	31092783          	lw	a5,784(s2)
   105fc:	008a16bb          	sllw	a3,s4,s0
   10600:	00892b83          	lw	s7,8(s2)
   10604:	8ff5                	and	a5,a5,a3
   10606:	2781                	sext.w	a5,a5
   10608:	eb99                	bnez	a5,1061e <__call_exitprocs+0xa6>
   1060a:	9702                	jalr	a4
   1060c:	00892783          	lw	a5,8(s2)
   10610:	f97799e3          	bne	a5,s7,105a2 <__call_exitprocs+0x2a>
   10614:	1f8c3783          	ld	a5,504(s8)
   10618:	fb2787e3          	beq	a5,s2,105c6 <__call_exitprocs+0x4e>
   1061c:	b759                	j	105a2 <__call_exitprocs+0x2a>
   1061e:	31492783          	lw	a5,788(s2)
   10622:	1084b583          	ld	a1,264(s1)
   10626:	8ff5                	and	a5,a5,a3
   10628:	2781                	sext.w	a5,a5
   1062a:	e799                	bnez	a5,10638 <__call_exitprocs+0xc0>
   1062c:	8556                	mv	a0,s5
   1062e:	9702                	jalr	a4
   10630:	bff1                	j	1060c <__call_exitprocs+0x94>
   10632:	00892423          	sw	s0,8(s2)
   10636:	b7c1                	j	105f6 <__call_exitprocs+0x7e>
   10638:	852e                	mv	a0,a1
   1063a:	9702                	jalr	a4
   1063c:	bfc1                	j	1060c <__call_exitprocs+0x94>

000000000001063e <_exit>:
   1063e:	4581                	li	a1,0
   10640:	4601                	li	a2,0
   10642:	4681                	li	a3,0
   10644:	4701                	li	a4,0
   10646:	4781                	li	a5,0
   10648:	05d00893          	li	a7,93
   1064c:	00000073          	ecall
   10650:	00054363          	bltz	a0,10656 <_exit+0x18>
   10654:	a001                	j	10654 <_exit+0x16>
   10656:	1141                	addi	sp,sp,-16
   10658:	e022                	sd	s0,0(sp)
   1065a:	842a                	mv	s0,a0
   1065c:	e406                	sd	ra,8(sp)
   1065e:	4080043b          	negw	s0,s0
   10662:	00000097          	auipc	ra,0x0
   10666:	00c080e7          	jalr	12(ra) # 1066e <__errno>
   1066a:	c100                	sw	s0,0(a0)
   1066c:	a001                	j	1066c <_exit+0x2e>

000000000001066e <__errno>:
   1066e:	00002797          	auipc	a5,0x2
   10672:	80a78793          	addi	a5,a5,-2038 # 11e78 <_impure_ptr>
   10676:	6388                	ld	a0,0(a5)
   10678:	8082                	ret
